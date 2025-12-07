package services

import (
	"context"
	"errors"
	"lelang-online-api/handlers"
	"lelang-online-api/models"
	"lelang-online-api/repositories"

	"gorm.io/gorm"
)

type AuctionService interface {
	NewAuction(ctx context.Context, auction *models.Auction) error
	All(ctx context.Context) ([]models.Auction, error)
	AddBid(ctx context.Context, userID uint, bid *models.BidRequest) error
}

type AuctionServiceImpl struct {
	repo repositories.AuctionRepo
}

func NewAuctionService(repo repositories.AuctionRepo) AuctionService {
	return &AuctionServiceImpl{repo: repo}
}

func (s *AuctionServiceImpl) NewAuction(ctx context.Context, auction *models.Auction) error {
	exist, err := s.repo.ItemExistAuction(ctx, auction)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("item sudah ada di lelang")
	}

	err = s.repo.Create(ctx, auction)
	if err != nil {
		return err
	}

	auctionNew, err := s.repo.FindByID(ctx, auction.ID)
	if err != nil {
		return err
	}

	handlers.BroadcastJson(auctionNew)

	return nil
}

func (s *AuctionServiceImpl) All(ctx context.Context) ([]models.Auction, error) {
	return s.repo.All(ctx)
}

func (s *AuctionServiceImpl) AddBid(ctx context.Context, userID uint, bid *models.BidRequest) error {
	exist, err := s.repo.ExistAndActivate(ctx, bid.ID)
	if err != nil {
		return err
	}

	if !exist {
		return gorm.ErrRecordNotFound
	}

	auction, err := s.repo.FindByID(ctx, bid.ID)
	if err != nil {
		return err
	}

	if auction.CurrentBid >= bid.Amount {
		return errors.New("harga penawaran tidak boleh lebih rendah atau sama dengan daripada harga saat ini")
	}

	history := models.AuctionHistory{
		AuctionID: auction.ID,
		ItemID:    auction.ItemID,
		UserID:    userID,
		Amount:    bid.Amount,
	}

	auction.BidCount = auction.BidCount + 1
	auction.CurrentBid = bid.Amount

	err = s.repo.UpdateBid(ctx, auction.ID, auction, history)
	if err != nil {
		return err
	}

	handlers.BroadcastJson(auction)

	return nil
}

package services

import (
	"context"
	"errors"
	"lelang-online-api/models"
	"lelang-online-api/repositories"
)

type AuctionService interface {
	NewAuction(ctx context.Context, auction *models.Auction) error
	All(ctx context.Context) ([]models.Auction, error)
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

	return s.repo.Create(ctx, auction)
}

func (s *AuctionServiceImpl) All(ctx context.Context) ([]models.Auction, error) {
	return s.repo.All(ctx)
}

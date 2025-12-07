package repositories

import (
	"context"
	"lelang-online-api/models"

	"gorm.io/gorm"
)

type AuctionRepo interface {
	ItemExistAuction(ctx context.Context, item *models.Auction) (bool, error)
	Create(ctx context.Context, item *models.Auction) error
	All(ctx context.Context) ([]models.Auction, error)
	ExistAndActivate(ctx context.Context, auctionID uint) (bool, error)
	UpdateBid(ctx context.Context, id uint, auction *models.Auction, history models.AuctionHistory) error
	FindByID(ctx context.Context, id uint) (*models.Auction, error)
}

type AuctionRepoImpl struct {
	DB *gorm.DB
}

func NewAuctionRepo(db *gorm.DB) AuctionRepo {
	return &AuctionRepoImpl{DB: db}
}

func (r *AuctionRepoImpl) All(ctx context.Context) ([]models.Auction, error) {
	var auctions []models.Auction
	err := r.DB.WithContext(ctx).Preload("Item").Preload("User").Preload("Histories").Find(&auctions).Error
	return auctions, err
}

func (r *AuctionRepoImpl) ItemExistAuction(ctx context.Context, item *models.Auction) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.Auction{}).
		Where("item_id = ? AND user_id = ?", item.ItemID, item.UserID).
		Count(&count).Error
	return count > 0, err
}

func (r *AuctionRepoImpl) Create(ctx context.Context, item *models.Auction) error {
	return r.DB.WithContext(ctx).Create(&item).Error
}

func (r *AuctionRepoImpl) ExistAndActivate(ctx context.Context, auctionID uint) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.Auction{}).
		Where("id = ?", auctionID).Where("status = 1").Count(&count).Error

	return count > 0, err
}

func (r *AuctionRepoImpl) UpdateBid(ctx context.Context, id uint, auction *models.Auction, history models.AuctionHistory) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.Auction{}).Where("id = ?", id).Updates(&auction).Error
		if err != nil {
			return err
		}

		if err := tx.Create(&history).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *AuctionRepoImpl) FindByID(ctx context.Context, id uint) (*models.Auction, error) {
	var auction models.Auction
	err := r.DB.WithContext(ctx).Find(&auction, "id = ?", id).Error
	return &auction, err
}

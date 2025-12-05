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
}

type AuctionRepoImpl struct {
	DB *gorm.DB
}

func NewAuctionRepo(db *gorm.DB) AuctionRepo {
	return &AuctionRepoImpl{DB: db}
}

func (r *AuctionRepoImpl) All(ctx context.Context) ([]models.Auction, error) {
	var auctions []models.Auction
	err := r.DB.WithContext(ctx).Preload("Item").Preload("User").Find(&auctions).Error
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

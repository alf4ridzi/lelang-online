package repositories

import (
	"context"
	"lelang-online-api/models"

	"gorm.io/gorm"
)

type ItemRepo interface {
	Create(context.Context, models.Item) error
	FindByID(context.Context, any) ([]models.Item, error)
}

type ItemRepoImpl struct {
	DB *gorm.DB
}

func NewItemRepo(DB *gorm.DB) ItemRepo {
	return &ItemRepoImpl{DB: DB}
}

func (r *ItemRepoImpl) Create(ctx context.Context, item models.Item) error {
	return r.DB.WithContext(ctx).Create(&item).Error
}

func (r *ItemRepoImpl) FindByID(ctx context.Context, id any) ([]models.Item, error) {
	var items []models.Item

	err := r.DB.WithContext(ctx).Preload("User").Find(&items, "user_id = ?", id).Error
	return items, err
}

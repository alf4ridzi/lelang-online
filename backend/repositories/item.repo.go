package repositories

import (
	"context"
	"lelang-online-api/models"

	"gorm.io/gorm"
)

type ItemRepo interface {
	Create(ctx context.Context, item models.Item) error
	FindByUserID(ctx context.Context, id any) ([]models.Item, error)
	FindByID(ctx context.Context, id any) (*models.Item, error)
	Exist(ctx context.Context, id any) (bool, error)
	Update(ctx context.Context, id any, item *models.Item) error
	Delete(ctx context.Context, id any) error
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

func (r *ItemRepoImpl) FindByUserID(ctx context.Context, id any) ([]models.Item, error) {
	var items []models.Item

	err := r.DB.WithContext(ctx).Preload("User").Find(&items, "user_id = ?", id).Error
	return items, err
}

func (r *ItemRepoImpl) FindByID(ctx context.Context, id any) (*models.Item, error) {
	var item models.Item
	err := r.DB.WithContext(ctx).Preload("User").Find(&item, "id = ?", id).Error
	return &item, err
}

func (r *ItemRepoImpl) Exist(ctx context.Context, id any) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.Item{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *ItemRepoImpl) Update(ctx context.Context, id any, item *models.Item) error {
	err := r.DB.WithContext(ctx).Model(&models.Item{}).Where("id = ?", id).Updates(&item).Error
	return err
}

func (r *ItemRepoImpl) Delete(ctx context.Context, id any) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&models.Item{}).Error
	return err
}

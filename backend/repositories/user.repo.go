package repositories

import (
	"context"
	"lelang-online-api/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindByUsername(context.Context, string) (*models.User, error)
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{DB: db}
}

func (r *UserRepoImpl) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).First(&user, "username = ?", username).Error
	return &user, err
}

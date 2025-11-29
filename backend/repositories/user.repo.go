package repositories

import "gorm.io/gorm"

type UserRepo interface{}

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{DB: db}
}

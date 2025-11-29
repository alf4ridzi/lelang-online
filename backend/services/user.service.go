package services

import (
	"context"
	"errors"
	"lelang-online-api/models"
	"lelang-online-api/repositories"
	"lelang-online-api/utils"

	"gorm.io/gorm"
)

type UserService interface {
	Login(context.Context, models.Login) (*models.User, error)
}

type UserServiceImpl struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Login(ctx context.Context, login models.Login) (*models.User, error) {
	user, err := s.repo.FindByUsername(ctx, login.Username)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("username atau password salah")
	}

	if err != nil {
		return nil, errors.New("internal server error")
	}

	if !utils.ValidateBcryptHash(login.Password, user.Password) {
		return nil, errors.New("username atau password salah")
	}

	return user, nil
}

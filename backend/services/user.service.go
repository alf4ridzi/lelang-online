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
	Login(ctx context.Context, login models.Login) (*models.User, error)
	Register(ctx context.Context, user models.User) error
	Profile(ctx context.Context, id any) (*models.User, error)
	GetItems(ctx context.Context, id any) ([]models.Item, error)
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

func (s *UserServiceImpl) Register(ctx context.Context, user models.User) error {
	hashed, err := utils.GenerateHashBcrypt(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed

	return s.repo.Create(ctx, user)
}

func (s *UserServiceImpl) Profile(ctx context.Context, id any) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserServiceImpl) GetItems(ctx context.Context, id any) ([]models.Item, error) {
	return s.repo.GetItems(ctx, id)
}

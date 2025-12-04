package services

import (
	"context"
	"errors"
	"lelang-online-api/models"
	"lelang-online-api/repositories"
)

type ItemService interface {
	Create(ctx context.Context, item models.Item) error
	GetByUserID(ctx context.Context, id any) ([]models.Item, error)
	Update(ctx context.Context, id any, item *models.Item) error
	Delete(ctx context.Context, id any) error
}

type ItemServiceImpl struct {
	repo repositories.ItemRepo
}

func NewItemService(repo repositories.ItemRepo) ItemService {
	return &ItemServiceImpl{repo: repo}
}

func (s *ItemServiceImpl) Create(ctx context.Context, item models.Item) error {
	return s.repo.Create(ctx, item)
}

func (s *ItemServiceImpl) GetByUserID(ctx context.Context, id any) ([]models.Item, error) {
	return s.repo.FindByUserID(ctx, id)
}

func (s *ItemServiceImpl) Update(ctx context.Context, id any, item *models.Item) error {
	exist, err := s.repo.Exist(ctx, id)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("record not found")
	}

	return s.repo.Update(ctx, id, item)
}

func (s *ItemServiceImpl) Delete(ctx context.Context, id any) error {
	return s.repo.Delete(ctx, id)
}

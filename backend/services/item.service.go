package services

import (
	"context"
	"lelang-online-api/models"
	"lelang-online-api/repositories"
)

type ItemService interface {
	Create(context.Context, models.Item) error
	GetByID(context.Context, any) ([]models.Item, error)
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

func (s *ItemServiceImpl) GetByID(ctx context.Context, id any) ([]models.Item, error) {
	return s.repo.FindByID(ctx, id)
}

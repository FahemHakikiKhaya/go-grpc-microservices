package service

import (
	"context"

	"github.com/fahemhakikikhaya/go-microservices-menu/internal/domain"
)

type menuService struct {
	store domain.MenuStore
}

// CreateMenu implements domain.MenuService.
func (m *menuService) CreateMenu(ctx context.Context, menu *domain.Menu) (*domain.Menu, error) {
	return m.store.Create(ctx, menu)
}

// GetMenuByID implements domain.MenuService.
func (m *menuService) GetMenuByID(ctx context.Context, id string) (*domain.Menu, error) {
	return m.store.GetByID(ctx, id)
}

func NewMenuService(store domain.MenuStore) domain.MenuService {
	return &menuService{
		store: store,
	}
}

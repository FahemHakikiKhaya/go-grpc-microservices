package service

import (
	"context"

	"github.com/fahemhakikikhaya/go-microservices-orders/internal/domain"
)

type orderService struct {
	store domain.OrderStore
}

// CreateOrder implements domain.OrderService.
func (o *orderService) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	return o.store.Create(ctx, order)
}

func NewOrderService(store domain.OrderStore) domain.OrderService {
	return &orderService{store}
}

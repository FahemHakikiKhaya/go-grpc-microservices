package store

import (
	"context"

	"github.com/fahemhakikikhaya/go-microservices-orders/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderStore struct {
	db *mongo.Collection
}

// Create implements OrderStore.
func (s *orderStore) Create(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	order.ID = primitive.NewObjectID().Hex()

	order.Status = "Pending"

	_, err := s.db.InsertOne(ctx, order)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func NewOrderStore(db *mongo.Collection) domain.OrderStore {
	return &orderStore{db}
}

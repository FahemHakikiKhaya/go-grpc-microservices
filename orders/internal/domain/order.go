package domain

import "context"

type Order struct {
	ID         string  `bson:"_id,omitempty" json:"id"`
	UserID     string  `bson:"userId" json:"userId"`
	Status     string  `bson:"status" json:"status"`
	TotalPrice float64 `bson:"totalPrice" json:"totalPrice"`
	Items []OrderItem  `bson:"items" json:"items"`
}

type OrderService interface {
	CreateOrder(context.Context, *Order) (*Order, error)
}

type OrderStore interface {
	Create(context.Context, *Order) (*Order, error)
}
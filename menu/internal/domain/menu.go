package domain

import (
	"context"
)

type Menu struct {
	ID string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Price float64 `bson:"price" json:"price"`
	Available bool `bson:"available" json:"available"`
}

type MenuStore interface {
	Create(ctx context.Context, menu *Menu) (*Menu, error)
	GetByID(ctx context.Context, id string) (*Menu, error)
}

type MenuService interface {
	CreateMenu(ctx context.Context, menu *Menu) (*Menu, error)
	GetMenuByID(ctx context.Context, id string) (*Menu, error)
}
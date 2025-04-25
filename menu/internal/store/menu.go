package store

import (
	"context"

	"github.com/fahemhakikikhaya/go-microservices-menu/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type menuStore struct {
	db *mongo.Collection
}

// Create implements domain.MenuStore.
func (m *menuStore) Create(ctx context.Context, menu *domain.Menu) (*domain.Menu, error) {
	menu.ID = primitive.NewObjectID().Hex()
	
	_, err := m.db.InsertOne(ctx, menu)

	if err != nil {
		return nil, err
	}

	return menu, nil
}

// GetByID implements domain.MenuStore.
func (m *menuStore) GetByID(ctx context.Context, id string) (*domain.Menu, error) {
	var menu domain.Menu

	err := m.db.FindOne(ctx, bson.M{"_id": id}).Decode(&menu)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	
	return &menu, nil
}

func NewMenuStore(db *mongo.Collection) domain.MenuStore {
	return &menuStore{
		db: db,
	}
}

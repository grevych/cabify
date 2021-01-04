package marketplace

import (
	"github.com/grevych/cabify/internal/storage"
	"github.com/grevych/cabify/pkg/entities"
)

type Marketplace struct {
	database *storage.Storage
}

func NewMarketplace(database *storage.Storage) *Marketplace {
	return &Marketplace{database}
}

func (m *Marketplace) ListProducts() ([]*entities.Product, error) {
	productStore := m.database.Products

	return productStore.All(), nil
}

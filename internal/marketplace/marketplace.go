package marketplace

import (
	"github.com/grevych/cabify/internal/storage/memory"
	"github.com/grevych/cabify/pkg/entities"
)

type Marketplace struct {
	storage *memory.Storage
}

func NewMarketplace(storage *memory.Storage) *Marketplace {
	return &Marketplace{storage}
}

func (m *Marketplace) ListProducts() ([]*entities.Product, error) {
	productStore := m.storage.GetProductStore()

	return productStore.All(), nil
}

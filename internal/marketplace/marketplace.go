package marketplace

import (
	stg "github.com/grevych/cabify/internal/storage"
	"github.com/grevych/cabify/pkg/entities"
)

type Marketplace struct {
	storage stg.Storage
}

func NewMarketplace(storage stg.Storage) *Marketplace {
	return &Marketpace{storage}
}

func (m *Marketplace) ListProducts() ([]*entities.Products, error) {
	products := []*entities.Products{}
	basketStore := m.storage.GetBasketStore()

	for _, entity := range basketStore.Items {
		product, err := entity.(Product)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

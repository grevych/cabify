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


func(m *Marketplace) ListProducts() ([]*entities.Products, error) {
  return []*entities.Product{}, nil
}

package storage

import (
	"github.com/grevych/cabify/pkg/entities"
)

type entityStore interface {
	Save(entities.Entity) (string, error)
	Update(entities.Entity) error
	Delete(string) error
}

type BasketStore interface {
	entityStore
	FindById(string) (*entities.Basket, error)
}

type ProductStore interface {
	entityStore
	FindById(string) (*entities.Product, error)
	All() []*entities.Product
}

type Storage struct {
	Baskets  BasketStore
	Products ProductStore
}

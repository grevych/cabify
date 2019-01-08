package storage

import (
	"github.com/grevych/cabify/pkg/entities"
)

type entityStore interface {
	Save(entities.Entity) (string, error)
	Update(entities.Entity) error
	Delete(string) error
}

type basketStore interface {
	entityStore
	FindById(string) (*entities.Basket, error)
}

type productStore interface {
	entityStore
	FindById(string) (*entities.Product, error)
}

type Storage interface {
	GetBasketStore() basketStore
	GetProductStore() productStore
}

package marketplace

import (
	"github.com/grevych/cabify/internal/storage"
	stg "github.com/grevych/cabify/internal/storage"
	"github.com/grevych/cabify/pkg/entities"
)

type Checkout struct {
	storage storage.Storage
}

func NewCheckout(storage stg.Storage) *Checkout {
	return &Checkout{storage}
}

func (c *Checkout) Detail(basketId string) (*entities.Basket, error) {
	return &entities.Basket{}, nil
}

func (c *Checkout) Delete(basketId string) error {
	return nil
}

func (c *Checkout) AddProduct(basketId string, product *entities.Product) error {
	return nil
}

func (c *Checkout) RemoveProduct(basketId string, product *entities.Product) error {
	return nil
}

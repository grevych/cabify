package marketplace

import (
	"github.com/grevych/cabify/internal/storage/memory"
	"github.com/grevych/cabify/pkg/entities"
)

type Checkout struct {
	storage *memory.Storage
}

func NewCheckout(storage *memory.Storage) *Checkout {
	return &Checkout{storage}
}

func (c *Checkout) Create() (*entities.Basket, error) {
	basket, _ := entities.NewBasket("", nil)
	basketStore := c.storage.GetBasketStore()
	_, err := basketStore.Save(basket)
	if err != nil {
		return nil, err
	}

	return basket, nil
}

func (c *Checkout) Detail(basketId string) (*entities.Basket, error) {
	basketStore := c.storage.GetBasketStore()
	return basketStore.FindById(basketId)
}

func (c *Checkout) Delete(basketId string) error {
	basketStore := c.storage.GetBasketStore()
	return basketStore.Delete(basketId)
}

func (c *Checkout) AddProduct(basketId, productId string) error {
	basketStore := c.storage.GetBasketStore()
	basket, err := basketStore.FindById(basketId)
	if err != nil {
		return err
	}

	productStore := c.storage.GetProductStore()
	product, err := productStore.FindById(productId)
	if err != nil {
		return err
	}

	if err := basket.AddProduct(product); err != nil {
		return err
	}

	if err := basketStore.Update(basket); err != nil {
		return err
	}

	return nil
}

func (c *Checkout) RemoveProduct(basketId, productId string) error {
	basketStore := c.storage.GetBasketStore()
	basket, err := basketStore.FindById(basketId)
	if err != nil {
		return err
	}

	if err := basket.RemoveProduct(productId); err != nil {
		return err
	}

	if err := basketStore.Update(basket); err != nil {
		return err
	}

	return nil
}

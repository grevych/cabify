package entities

import (
	"errors"
	"fmt"
)

type Basket struct {
	*entity
	Products []*Product
}

var _ Entity = &Basket{}

func NewBasket(id string, products []*Product) (*Basket, error) {
	if products == nil {
		products = []*Product{}
	}

	return &Basket{
		entity:   &entity{id},
		Products: products,
	}, nil
}

func (b *Basket) AddProduct(product *Product) error {
	if product == nil {
		return errors.New("Product cannot be nil")
	}

	b.Products = append(b.Products, product)

	return nil
}

func (b *Basket) RemoveProduct(productId string) error {
	products := []*Product{}
	productNotFound := true

	for _, product := range b.Products {
		if product.GetId() == productId && productNotFound {
			productNotFound = false
			continue
		}

		products = append(products, product)
	}

	if len(products) == len(b.Products) {
		return fmt.Errorf("Product %s does not exist in basket %s", productId, b.GetId())
	}

	b.Products = products

	return nil
}

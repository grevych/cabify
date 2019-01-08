package entities

import (
  "errors"
  "fmt"
)

type Basket struct {
  *entity
  Products map[string]*Product
}

func NewBasket(id string, products map[string]*Product) (*Basket, error) {
	return &Basket{
		entity: &entity{id},
		Products: products,
	}, nil
}

func(b *Basket) AddProduct(product *Product) error {
	if product == nil {
		return errors.New("Product cannot be nil")
	}	

	if product.id == "" {
		return errors.New("Product requires an id")
	}

	if _, ok := b.Products[product.id]; ok {
		return fmt.Errorf("Product %s already exists in basket %s", product.id, b.id)
	}

	b.Products[product.id] = product
	return nil
}

func(b *Basket) RemoveProduct(productId string) error {
	if _, ok := b.Products[productId]; !ok {
		return fmt.Errorf("Product %s does not exist in basket %s", productId, b.id)
	}

	delete(b.Products, productId)
	return nil
}

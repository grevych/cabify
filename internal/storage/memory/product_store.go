package memory

import (
	"fmt"

	"github.com/grevych/cabify/pkg/entities"
)

type ProductStore struct {
	*store
}

func NewProductStore() *ProductStore {
	return &ProductStore{
		store: NewStore(),
	}
}

func (ps *ProductStore) All() []*entities.Product {
	products := []*entities.Product{}

	for _, entity := range ps.store.All() {
		if product, ok := entity.(*entities.Product); ok {
			products = append(products, product)
		}
	}

	return products
}

func (ps *ProductStore) FindById(productId string) (*entities.Product, error) {
	entity, err := ps.store.FindById(productId)

	if err != nil {
		return nil, err
	}

	product, ok := entity.(*entities.Product)

	if !ok {
		return nil, fmt.Errorf("Invalid entity type for Product %s", productId)
	}

	return product, nil
}

package memory

import (
	"fmt"

	"github.com/grevych/cabify/pkg/entities"
)

type ProductStore struct {
	*store
}

func NewProductStore() *ProductStore {
	ps := &ProductStore{
		store: NewStore(),
	}

	// Add Voucher
	voucher, _ := entities.NewProduct("", "VOUCHER", "Cabify Voucher", 400)
	ps.Save(voucher)

	// Add Mug
	mug, _ := entities.NewProduct("", "MUG", "Cabify Mug", 500)
	ps.Save(mug)

	// Add Shirt
	shirt, _ := entities.NewProduct("", "SHIRT", "Cabify Shirt", 600)
	ps.Save(shirt)

	return ps
}

func (ps *ProductStore) FindById(productId string) (*entities.Product, error) {
	obj, err := ps.store.FindEntityById(productId)

	if err != nil {
		return nil, err
	}

	product, ok := obj.(entities.Product)

	if !ok {
		return nil, fmt.Errorf("Unable to assert type Product for id %s", productId)
	}

	return &product, nil
}

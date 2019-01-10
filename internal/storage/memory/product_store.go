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

	// Add Voucher
	// voucher, _ := entities.NewProduct("", "VOUCHER", "Cabify Voucher", 400)
	// ps.store.Save(voucher)

	// Add Mug
	// mug, _ := entities.NewProduct("", "MUG", "Cabify Mug", 500)
	// ps.store.Save(mug)

	// Add Shirt
	// shirt, _ := entities.NewProduct("", "SHIRT", "Cabify Shirt", 600)
	// ps.store.Save(shirt)
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

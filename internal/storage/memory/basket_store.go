package memory

import (
	"fmt"

	"github.com/grevych/cabify/pkg/entities"
)

type BasketStore struct {
	*store
}

func NewBasketStore() *BasketStore {
	return &BasketStore{
		store: NewStore(),
	}
}

func (bs *BasketStore) FindById(basketId string) (*entities.Basket, error) {
	entity, err := bs.store.FindById(basketId)

	if err != nil {
		return nil, err
	}

	basket, ok := entity.(*entities.Basket)

	if !ok {
		return nil, fmt.Errorf("Invalid entity type for Basket %s", basketId)
	}

	return basket, nil
}

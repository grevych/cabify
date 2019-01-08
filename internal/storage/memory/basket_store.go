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
	obj, err := bs.store.FindEntityById(basketId)

	if err != nil {
		return nil, err
	}

	basket, ok := obj.(entities.Basket)

	if !ok {
		return nil, fmt.Errorf("Unable to assert type Basket for id %s", basketId)
	}

	return &basket, nil
}

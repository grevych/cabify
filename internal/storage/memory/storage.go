package memory

import "errors"

type Storage struct {
	BasketStore  *BasketStore
	ProductStore *ProductStore
}

func NewStorage(bs *BasketStore, ps *ProductStore) (*Storage, error) {
	if bs == nil {
		return nil, errors.New("Basket store cannot be nil")
	}

	if ps == nil {
		return nil, errors.New("Product store cannot be nil")
	}

	return &Storage{
		BasketStore:  bs,
		ProductStore: ps,
	}, nil
}

func (s *Storage) GetBasketStore() *BasketStore {
	return s.BasketStore
}

func (s *Storage) GetProductStore() *ProductStore {
	return s.ProductStore
}

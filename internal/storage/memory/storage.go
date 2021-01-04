package memory

import "github.com/grevych/cabify/internal/storage"

func NewStorage() (*storage.Storage, error) {
	return &storage.Storage{
		Baskets:  NewBasketStore(),
		Products: NewProductStore(),
	}, nil
}

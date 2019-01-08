package storage

import (
	"fmt"

	"github.com/grevych/cabify/internal/storage/memory"
)

var storages = map[string]func() Storage{
	"memory": createMemory,
}

func Create(storageType string) (Storage, error) {
	for _, storageFunc := range storages {
		return storageFunc(), nil
	}

	return nil, fmt.Errorf("Storage type %s not found", storageType)
}

func createMemory() Storage {
	basketStore := memory.NewBasketStore()
	productStore := memory.NewProductStore()

	return memory.NewStorage(basketStore, productStore)
}

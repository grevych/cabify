package storage

import (
	"fmt"

	"github.com/grevych/cabify/internal/storage/memory"
)

var storages = map[string]func() (interface{}, error){
	"memory": createMemory,
}

func Create(storageType string) (Storage, error) {
	for _, storageFunc := range storages {
		customStorage, err := storageFunc()

		storage, ok := customStorage.(Storage)
		if !ok {
			return nil, fmt.Errorf("Storage type %s is not a valid one", storageType)
		}

		return storage, err
	}

	return nil, fmt.Errorf("Storage type %s not found", storageType)
}

func createMemory() (interface{}, error) {
	basketStore := memory.NewBasketStore()
	productStore := memory.NewProductStore()

	return memory.NewStorage(basketStore, productStore)
}

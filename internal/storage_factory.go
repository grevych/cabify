package internal

import (
	"fmt"

	"github.com/grevych/cabify/internal/storage"
	"github.com/grevych/cabify/internal/storage/memory"
)

// Create is a factory that initializes storages on demand
func CreateStorage(storageName string) (*storage.Storage, error) {
	switch storageName {
	case "memory":
		return memory.NewStorage()
	default:
		return nil, fmt.Errorf("Storage '%s' not found", storageName)
	}
}

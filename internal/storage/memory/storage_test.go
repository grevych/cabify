package memory

import (
	"reflect"
	"testing"

	"github.com/grevych/cabify/internal/storage"
)

func createStorage() *storage.Storage {
	storage, _ := NewStorage()
	return storage
}

func TestNewStorage(t *testing.T) {
	tests := map[string]struct {
		expectedStorage *storage.Storage
		expectedError   error
	}{
		"TestNewStorage": {
			expectedStorage: &storage.Storage{
				Baskets:  NewBasketStore(),
				Products: NewProductStore(),
			},
			expectedError: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			storage, err := NewStorage()

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			if !reflect.DeepEqual(testCase.expectedStorage, storage) {
				t.Errorf(
					"Expected storage %+v, got %+v",
					testCase.expectedStorage,
					storage,
				)
			}
		})
	}
}

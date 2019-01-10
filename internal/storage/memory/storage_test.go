package memory

import (
	"errors"
	"reflect"
	"testing"
)

func createStorage(bs *BasketStore, ps *ProductStore) *Storage {
	storage, _ := NewStorage(bs, ps)
	return storage
}

func TestNewStorage(t *testing.T) {
	tests := map[string]struct {
		basketStore     *BasketStore
		productStore    *ProductStore
		expectedStorage *Storage
		expectedError   error
	}{
		"TestNewStorage": {
			basketStore:  NewBasketStore(),
			productStore: NewProductStore(),
			expectedStorage: &Storage{
				BasketStore:  NewBasketStore(),
				ProductStore: NewProductStore(),
			},
			expectedError: nil,
		},

		"TestNewStorageWithNilBasketStore": {
			basketStore:     nil,
			productStore:    NewProductStore(),
			expectedStorage: nil,
			expectedError:   errors.New("Basket store cannot be nil"),
		},

		"TestNewStorageWithNilProductStore": {
			basketStore:     NewBasketStore(),
			productStore:    nil,
			expectedStorage: nil,
			expectedError:   errors.New("Product store cannot be nil"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			storage, err := NewStorage(
				testCase.basketStore, testCase.productStore,
			)

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

func TestGetBasketStore(t *testing.T) {
	basketStore := NewBasketStore()
	productStore := NewProductStore()

	tests := map[string]struct {
		storage             *Storage
		expectedBasketStore *BasketStore
	}{
		"TestGetBasketStore": {
			storage:             createStorage(basketStore, productStore),
			expectedBasketStore: basketStore,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			basketStore := testCase.storage.GetBasketStore()

			if !reflect.DeepEqual(testCase.expectedBasketStore, basketStore) {
				t.Errorf(
					"Expected basket store %+v, got %+v",
					testCase.expectedBasketStore,
					basketStore,
				)
			}
		})
	}
}

func TestGetProductStore(t *testing.T) {
	basketStore := NewBasketStore()
	productStore := NewProductStore()

	tests := map[string]struct {
		storage              *Storage
		expectedProductStore *ProductStore
	}{
		"TestGetProductStore": {
			storage:              createStorage(basketStore, productStore),
			expectedProductStore: productStore,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			productStore := testCase.storage.GetProductStore()

			if !reflect.DeepEqual(testCase.expectedProductStore, productStore) {
				t.Errorf(
					"Expected product store %+v, got %+v",
					testCase.expectedProductStore,
					productStore,
				)
			}
		})
	}
}

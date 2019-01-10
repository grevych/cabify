package memory

import (
	"errors"
	"reflect"
	"testing"

	"github.com/grevych/cabify/pkg/entities"
)

func TestNewBasketStore(t *testing.T) {
	tests := map[string]struct {
		expectedBasketStore *BasketStore
	}{
		"TestNewBasketStore": {
			expectedBasketStore: &BasketStore{
				store: NewStore(),
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			basketStore := NewBasketStore()

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

func TestFindByBasketId(t *testing.T) {
	basketId := "basket-id"
	product, _ := entities.NewProduct("", "PRODUCT", "product", 100)

	tests := map[string]struct {
		basketId       string
		basketStore    *BasketStore
		expectedBasket *entities.Basket
		expectedError  error
	}{
		"TestFindById": {
			basketId:       basketId,
			basketStore:    createBasketStore(createBasket(basketId, product)),
			expectedBasket: createBasket(basketId, product),
			expectedError:  nil,
		},

		"TestFindByIdNonExistentBasket": {
			basketId:       "non-existent-basket-id",
			basketStore:    createBasketStore(),
			expectedBasket: nil,
			expectedError:  errors.New("Entity non-existent-basket-id not found"),
		},

		"TestFindByIdInvalidBasket": {
			basketId: basketId,
			basketStore: createBasketStore(
				createProduct("basket-id", "PRODUCT", "product", 500),
			),
			expectedBasket: nil,
			expectedError:  errors.New("Invalid entity type for Basket basket-id"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			basket, err := testCase.basketStore.FindById(testCase.basketId)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}

			if !reflect.DeepEqual(testCase.expectedBasket, basket) {
				t.Errorf("Expected basket %+v, got %+v", testCase.expectedBasket, basket)
			}
		})
	}
}

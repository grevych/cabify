package memory

import (
	"errors"
	"reflect"
	"testing"

	"github.com/grevych/cabify/pkg/entities"
)

func TestNewProductStore(t *testing.T) {
	tests := map[string]struct {
		expectedProductStore *ProductStore
	}{
		"TestNewProductStore": {
			expectedProductStore: &ProductStore{
				store: NewStore(),
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			productStore := NewProductStore()

			if !reflect.DeepEqual(testCase.expectedProductStore, productStore) {
				t.Errorf(
					"Expected basket store %+v, got %+v",
					testCase.expectedProductStore,
					productStore,
				)
			}
		})
	}
}

func TestFindByProductId(t *testing.T) {
	var productPrice int64 = 5000

	productId := "product-id"
	productCode := "PRODUCT"
	productName := "product"

	tests := map[string]struct {
		productId       string
		productStore    *ProductStore
		expectedProduct *entities.Product
		expectedError   error
	}{
		"TestFindById": {
			productId: productId,
			productStore: createProductStore(
				createProduct(productId, productCode, productName, productPrice),
			),
			expectedProduct: createProduct(
				productId, productCode, productName, productPrice,
			),
			expectedError: nil,
		},

		"TestFindByIdNonExistentProduct": {
			productId:       "non-existent-product-id",
			productStore:    createProductStore(),
			expectedProduct: nil,
			expectedError:   errors.New("Entity non-existent-product-id not found"),
		},

		"TestFindByIdInvalidBasket": {
			productId:       productId,
			productStore:    createProductStore(createBasket("product-id", nil)),
			expectedProduct: nil,
			expectedError:   errors.New("Invalid entity type for Product product-id"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			product, err := testCase.productStore.FindById(testCase.productId)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}

			if !reflect.DeepEqual(testCase.expectedProduct, product) {
				t.Errorf("Expected product %+v, got %+v", testCase.expectedProduct, product)
			}
		})
	}
}

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

func TestAllProducts(t *testing.T) {
	productA, _ := entities.NewProduct("product-A-id", "PRODUCT", "product", 50)
	productB, _ := entities.NewProduct("product-B-id", "PRODUCT", "product", 60)

	tests := map[string]struct {
		productStore     *ProductStore
		expectedProducts []*entities.Product
	}{
		"TestAll": {
			productStore:     createProductStore(productA, productB),
			expectedProducts: []*entities.Product{productA, productB},
		},

		"TestAllWithEmptyProductStore": {
			productStore:     createProductStore(),
			expectedProducts: []*entities.Product{},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			products := testCase.productStore.All()

			if !reflect.DeepEqual(testCase.expectedProducts, products) {
				t.Errorf(
					"Expected products %+v, got %+v",
					testCase.expectedProducts,
					products,
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

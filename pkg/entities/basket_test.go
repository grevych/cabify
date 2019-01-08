package entities

import (
	"errors"
	"testing"
	"reflect"
)


func createBasket(basketId string, maps ...map[string]*Product) *Basket {
	products := map[string]*Product{}

	if len(maps) > 0 {
		products = maps[0]  
	}

	basket, _ := NewBasket(basketId, products)
	return basket
}


func TestNewBasket(t *testing.T) {
	product, _ := NewProduct("product-id", "PRODUCT", "product", 50000)

	tests := map[string]struct {
		basketId string
		products map[string]*Product
		expectedBasket *Basket
		expectedError error
	}{
		"TestNewBasket": {
			basketId: "basket-id",
			products: map[string]*Product{
			  product.id: product,
			},
			expectedBasket: &Basket{
			  entity: &entity{"basket-id"},
			  Products: map[string]*Product{
				product.id: product,
			  },
			},
			expectedError: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			basket, err := NewBasket(testCase.basketId, testCase.products)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}
			
			if !reflect.DeepEqual(testCase.expectedBasket, basket) {
				t.Errorf("Expected basket %+v, got %+v", testCase.expectedBasket, basket)
			}
		})
	}
}


func TestAddProduct(t *testing.T) {
	product, _ := NewProduct("product-id", "PRODUCT", "product", 5000)

	tests := map[string]struct {
		basket *Basket
		product *Product
		expectedProducts map[string]*Product
		expectedError error
	}{
		"TestAddProduct": {
			basket: createBasket("basket-id"),
			product: product,
			expectedProducts: map[string]*Product{
				product.id: product,
			},
			expectedError: nil,
		},

		"TestAddProductTwice": {
			basket: createBasket("basket-id", map[string]*Product{
				product.id: product,
			}),
			product: product,
			expectedProducts: map[string]*Product{
				product.id: product,
			},
			expectedError: errors.New("Product product-id already exists in basket basket-id"),
		},

		"TestAddProductWithMissingId": {
			basket: createBasket("basket-id"),
			product: &Product{
			  entity: &entity{""},
			  Code: "PRODUCT",
			  Name: "product",
			  Price: 50000,
			},
			expectedProducts: map[string]*Product{},
			expectedError: errors.New("Product requires an id"),
		},

		"TestAddNilProduct": {
			basket: createBasket("basket-id", map[string]*Product{}),
			product: nil,
			expectedProducts: map[string]*Product{},
			expectedError: errors.New("Product cannot be nil"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.basket.AddProduct(testCase.product)
			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}
			
			if !reflect.DeepEqual(testCase.expectedProducts, testCase.basket.Products) {
				t.Errorf("Expected product %+v, got %+v", testCase.expectedProducts, testCase.basket.Products)
			}
		})
	}
}


func TestRemoveProductProduct(t *testing.T) {
	product, _ := NewProduct("product-id", "PRODUCT", "product", 5000)

	tests := map[string]struct {
		basket *Basket
		productId string
		expectedProducts map[string]*Product
		expectedError error
	}{
		"TestRemoveProduct": {
			basket: createBasket("basket-id", map[string]*Product{
				product.id: product,
			}),
			productId: product.id,
			expectedProducts: map[string]*Product{},
			expectedError: nil,
		},

		"TestRemoveNonExistentProduct": {
			basket: createBasket("basket-id", map[string]*Product{}),
			productId: product.id,
			expectedProducts: map[string]*Product{},
			expectedError: errors.New("Product product-id does not exist in basket basket-id"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.basket.RemoveProduct(testCase.productId)
			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}
			
			if !reflect.DeepEqual(testCase.expectedProducts, testCase.basket.Products) {
				t.Errorf("Expected product %+v, got %+v", testCase.expectedProducts, testCase.basket.Products)
			}
		})
	}
}

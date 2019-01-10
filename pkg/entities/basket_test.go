package entities

import (
	"errors"
	"reflect"
	"testing"
)

func createBasket(basketId string, products ...*Product) *Basket {
	basket, _ := NewBasket(basketId, products)
	return basket
}

func TestNewBasket(t *testing.T) {
	basketId := "basket-id"
	product, _ := NewProduct("product-id", "PRODUCT", "product", 50000)

	tests := map[string]struct {
		basketId       string
		products       []*Product
		expectedBasket *Basket
		expectedError  error
	}{
		"TestNewBasket": {
			basketId: basketId,
			products: []*Product{product},
			expectedBasket: &Basket{
				entity:   &entity{basketId},
				Products: []*Product{product},
			},
			expectedError: nil,
		},

		"TestNewBasketWithNilProducts": {
			basketId: basketId,
			products: nil,
			expectedBasket: &Basket{
				entity:   &entity{basketId},
				Products: []*Product{},
			},
			expectedError: nil,
		},

		"TestNewBasketWithEmptyId": {
			basketId: "",
			products: []*Product{product},
			expectedBasket: &Basket{
				entity:   &entity{},
				Products: []*Product{product},
			},
			expectedError: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			basket, err := NewBasket(testCase.basketId, testCase.products)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			if !reflect.DeepEqual(testCase.expectedBasket, basket) {
				t.Errorf(
					"Expected basket %+v, got %+v",
					testCase.expectedBasket,
					basket,
				)
			}
		})
	}
}

func TestAddProduct(t *testing.T) {
	basketId := "basket-id"
	product, _ := NewProduct("product-id", "PRODUCT", "product", 5000)

	tests := map[string]struct {
		basket         *Basket
		product        *Product
		expectedBasket *Basket
		expectedError  error
	}{
		"TestAddProduct": {
			basket:         createBasket(basketId),
			product:        product,
			expectedBasket: createBasket(basketId, product),
			expectedError:  nil,
		},

		"TestAddNilProduct": {
			basket:         createBasket(basketId),
			product:        nil,
			expectedBasket: createBasket(basketId),
			expectedError:  errors.New("Product cannot be nil"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.basket.AddProduct(testCase.product)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			if !reflect.DeepEqual(testCase.expectedBasket, testCase.basket) {
				t.Errorf(
					"Expected basket %+v, got %+v",
					testCase.expectedBasket,
					testCase.basket,
				)
			}
		})
	}
}

func TestRemoveProductProduct(t *testing.T) {
	basketId := "basket-id"
	product, _ := NewProduct("product-id", "PRODUCT", "product", 5000)

	tests := map[string]struct {
		basket         *Basket
		productId      string
		expectedBasket *Basket
		expectedError  error
	}{
		"TestRemoveProduct": {
			basket:         createBasket(basketId, product),
			productId:      product.GetId(),
			expectedBasket: createBasket(basketId),
			expectedError:  nil,
		},

		"TestRemoveNonExistentProduct": {
			basket:         createBasket(basketId),
			productId:      product.GetId(),
			expectedBasket: createBasket(basketId),
			expectedError: errors.New(
				"Product product-id does not exist in basket basket-id",
			),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := testCase.basket.RemoveProduct(testCase.productId)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			if !reflect.DeepEqual(testCase.expectedBasket, testCase.basket) {
				t.Errorf(
					"Expected product %+v, got %+v",
					testCase.expectedBasket,
					testCase.basket,
				)
			}
		})
	}
}

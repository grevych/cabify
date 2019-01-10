package entities

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {
	var productPrice int64 = 50000

	productId := "product-id"
	productCode := "PRODUCT"
	productName := "product"

	tests := map[string]struct {
		productId       string
		productCode     string
		productName     string
		productPrice    int64
		expectedProduct *Product
		expectedError   error
	}{
		"TestNewProduct": {
			productId:    productId,
			productCode:  productCode,
			productName:  productName,
			productPrice: productPrice,
			expectedProduct: &Product{
				entity: &entity{productId},
				Code:   productCode,
				Name:   productName,
				Price:  productPrice,
			},
			expectedError: nil,
		},

		"TestNewProductWithNegativePrice": {
			productId:       productId,
			productCode:     productCode,
			productName:     productName,
			productPrice:    -1000,
			expectedProduct: nil,
			expectedError:   errors.New("Product product-id requires a positive price"),
		},

		"TestNewProductWithZeroAsPrice": {
			productId:       productId,
			productCode:     productCode,
			productName:     productName,
			productPrice:    0,
			expectedProduct: nil,
			expectedError:   errors.New("Product product-id requires a positive price"),
		},

		"TestNewProductWithNoCode": {
			productId:       productId,
			productCode:     "",
			productName:     productName,
			productPrice:    productPrice,
			expectedProduct: nil,
			expectedError:   errors.New("Product product-id requires a code"),
		},

		"TestNewProductWithNoName": {
			productId:       productId,
			productCode:     productCode,
			productName:     "",
			productPrice:    productPrice,
			expectedProduct: nil,
			expectedError:   errors.New("Product product-id requires a name"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			product, err := NewProduct(
				testCase.productId,
				testCase.productCode,
				testCase.productName,
				testCase.productPrice,
			)

			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf(
					"Expected error %+v, got %+v",
					testCase.expectedError,
					err,
				)
			}

			if !reflect.DeepEqual(testCase.expectedProduct, product) {
				t.Errorf(
					"Expected product %+v, got %+v",
					testCase.expectedProduct,
					product,
				)
			}
		})
	}
}

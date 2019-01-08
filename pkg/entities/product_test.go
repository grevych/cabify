package entities

import (
	"errors"
	"testing"
	"reflect"
)


func TestNewProduct(t *testing.T) {
	tests := map[string]struct {
		productId string
		productCode string
		productName string
		productPrice int64
		expectedProduct *Product
		expectedError error
	}{
		"TestNewProduct": {
			productId: "product-id",
			productCode: "PRODUCT",
			productName: "product",
			productPrice: 50000,
			expectedProduct: &Product{
			  entity: &entity{"product-id"},
			  Code: "PRODUCT",
			  Name: "product",
			  Price: 50000,
			},
			expectedError: nil,
		},

		"TestNewProductWithNegativePrice": {
			productId: "product-id",
			productCode: "PRODUCT",
			productName: "product",
			productPrice: -1000,
			expectedProduct: nil,
			expectedError: errors.New("Product product-id requires a positive price"),
		},

		"TestNewProductWithZeroAsPrice": {
			productId: "product-id",
			productCode: "PRODUCT",
			productName: "product",
			productPrice: 0,
			expectedProduct: nil,
			expectedError: errors.New("Product product-id requires a positive price"),
		},

		"TestNewProductWithNoCode": {
			productId: "product-id",
			productCode: "",
			productName: "product",
			productPrice: 50000,
			expectedProduct: nil,
			expectedError: errors.New("Product product-id requires a code"),
		},

		"TestNewProductWithNoName": {
			productId: "product-id",
			productCode: "PRODUCT",
			productName: "",
			productPrice: 50000,
			expectedProduct: nil,
			expectedError: errors.New("Product product-id requires a name"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			product, err := NewProduct(
			  testCase.productId, testCase.productCode, testCase.productName, testCase.productPrice,
			)
			if !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Expected error %+v, got %+v", testCase.expectedError, err)
			}
			
			if !reflect.DeepEqual(testCase.expectedProduct, product) {
				t.Errorf("Expected product %+v, got %+v", testCase.expectedProduct, product)
			}
		})
	}
}

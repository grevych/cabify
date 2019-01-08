package entities

import (
  "fmt"
)


type Product struct {
  *entity
  Code string
  Name string

  // Price will be managed in cents
  Price int64
}

func NewProduct(id, code, name string, price int64) (*Product, error) {
	if price <= 0 {
		return nil, fmt.Errorf("Product %s requires a positive price", id)
	}
	
	if code == "" {
		return nil, fmt.Errorf("Product %s requires a code", id)
	}

	if name == "" {
		return nil, fmt.Errorf("Product %s requires a name", id)
	}

	return &Product{
		entity: &entity{id},
		Code: code,
		Name: name,
		Price: price,
	}, nil
}

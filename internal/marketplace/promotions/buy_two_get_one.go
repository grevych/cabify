package promotions

import "github.com/grevych/cabify/pkg/entities"

func PayTwoGetOneFree(productId string) Promotion {
	return func(basket *entities.Basket) {
		var productTarget *entities.Product

		count := 0
		for _, product := range basket.Products {
			if product.GetId() == productId {
				count += 1
				productTarget = product
			}
		}

		extras := count / 2
		for ; extras > 0; extras-- {
			newProduct, _ := entities.NewProduct(
				productTarget.Id,
				productTarget.Code,
				productTarget.Name,
				productTarget.Price,
			)
			newProduct.Price = 0
			basket.AddProduct(newProduct)
		}
	}
}

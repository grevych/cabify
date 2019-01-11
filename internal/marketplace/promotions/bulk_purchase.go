package promotions

import "github.com/grevych/cabify/pkg/entities"

func BulkPurchase(productId string, newPrice int64) Promotion {
	return func(basket *entities.Basket) {
		var (
			productTargetOne   *entities.Product
			productTargetTwo   *entities.Product
			productTargetThree *entities.Product
			oldPrice           int64
		)

		for _, product := range basket.Products {
			if product.GetId() != productId {
				continue
			}

			if productTargetOne == nil {
				productTargetOne = product
				oldPrice = product.Price
				continue
			}

			if productTargetTwo == nil {
				productTargetTwo = product
				continue
			}

			if productTargetThree == nil {
				productTargetThree = product
			}

			product.Price = newPrice
		}

		if productTargetThree == nil {
			if productTargetOne != nil {
				productTargetOne.Price = oldPrice
			}

			if productTargetTwo != nil {
				productTargetTwo.Price = oldPrice
			}
		}
	}
}

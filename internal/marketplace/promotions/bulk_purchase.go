package promotions

import "github.com/grevych/cabify/pkg/entities"

func BulkPurchase(productId string, newPrice int64) Promotion {
	return func(basket *entities.Basket) {
		var (
			productTargetOne   *entities.Product
			productTargetTwo   *entities.Product
			productTargetThree *entities.Product
		)

		for _, product := range basket.Products {
			if product.GetId() != productId {
				continue
			}

			if productTargetOne == nil {
				productTargetOne = product
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

		if productTargetThree != nil {
			productTargetOne.Price = newPrice
			productTargetTwo.Price = newPrice
		}
	}
}

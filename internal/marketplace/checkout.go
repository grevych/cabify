package marketplace

import (
	"errors"
	"log"

	"github.com/grevych/cabify/internal/marketplace/promotions"
	"github.com/grevych/cabify/internal/storage"
	"github.com/grevych/cabify/pkg/entities"
)

type Checkout struct {
	database   *storage.Storage
	promotions []promotions.Promotion
}

func NewCheckout(database *storage.Storage, initialPromotions []promotions.Promotion) *Checkout {
	return &Checkout{database, initialPromotions}
}

func (c *Checkout) Create() (*entities.Basket, error) {
	basket, _ := entities.NewBasket("", nil)
	basketStore := c.database.Baskets

	_, err := basketStore.Save(basket)
	if err != nil {
		message := "Error in checkout.Create while saving basket"
		log.Printf("%s: %+v", message, err)
		return nil, NewNotCreatedError(message)
	}

	return basket, nil
}

func (c *Checkout) Detail(basketId string) (*entities.Basket, error) {
	basketStore := c.database.Baskets

	basket, err := basketStore.FindById(basketId)
	if err != nil {
		message := "Error in checkout.Detail while finding basket"
		log.Printf("%s: %+v", message, err)
		return nil, NewNotFoundError(message)
	}

	basket = cloneBasket(basket)

	for _, promotion := range c.promotions {
		promotion(basket)
	}

	for _, product := range basket.Products {
		basket.Total += product.Price
	}

	return basket, nil
}

func (c *Checkout) Delete(basketId string) error {
	basketStore := c.database.Baskets

	if err := basketStore.Delete(basketId); err != nil {
		message := "Error in checkout.Delete while deleting basket"
		log.Printf("%s: %+v", message, err)
		return NewNotDeletedError(message)
	}

	return nil
}

func (c *Checkout) AddProduct(basketId, productId string) error {
	basketStore := c.database.Baskets
	productStore := c.database.Products

	basket, err := basketStore.FindById(basketId)
	if err != nil {
		message := "Error in checkout.AddProduct while finding basket"
		log.Printf("%s: %+v", message, err)
		return NewNotFoundError(message)
	}

	product, err := productStore.FindById(productId)
	if err != nil {
		message := "Error in checkout.AddProduct while finding product"
		log.Printf("%s: %+v", message, err)
		return NewNotFoundError(message)
	}

	product = cloneProduct(product)

	if err := basket.AddProduct(product); err != nil {
		message := "Error in checkout.AddProduct while adding product"
		log.Printf("%s: %+v", message, err)
		return errors.New(message)
	}

	if err := basketStore.Update(basket); err != nil {
		message := "Error in checkout.AddProduct while updating basket"
		log.Printf("%s: %+v", message, err)
		return NewNotUpdatedError(message)
	}

	return nil
}

func (c *Checkout) RemoveProduct(basketId, productId string) error {
	basketStore := c.database.Baskets

	basket, err := basketStore.FindById(basketId)
	if err != nil {
		message := "Error in checkout.RemoveProduct while finding basket"
		log.Printf("%s: %+v", message, err)
		return NewNotFoundError(message)
	}

	if err := basket.RemoveProduct(productId); err != nil {
		message := "Error in checkout.RemoveProduct while finding product"
		log.Printf("%s: %+v", message, err)
		return NewNotFoundError(message)
	}

	if err := basketStore.Update(basket); err != nil {
		message := "Error in checkout.RemoveProduct while updating basket"
		log.Printf("%s: %+v", message, err)
		return NewNotUpdatedError(message)
	}

	return nil
}

func cloneBasket(oldBasket *entities.Basket) *entities.Basket {
	newProducts := []*entities.Product{}

	for _, oldProduct := range oldBasket.Products {
		newProduct := cloneProduct(oldProduct)
		newProducts = append(newProducts, newProduct)
	}

	newBasket, _ := entities.NewBasket(oldBasket.Id, newProducts)

	return newBasket
}

func cloneProduct(oldProduct *entities.Product) *entities.Product {
	newProduct, _ := entities.NewProduct(
		oldProduct.Id,
		oldProduct.Code,
		oldProduct.Name,
		oldProduct.Price,
	)

	return newProduct
}

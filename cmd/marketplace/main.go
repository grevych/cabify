package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/grevych/cabify/internal"
	"github.com/grevych/cabify/internal/handlers"
	"github.com/grevych/cabify/internal/marketplace"
	"github.com/grevych/cabify/internal/marketplace/promotions"
	"github.com/grevych/cabify/internal/storage"
	"github.com/grevych/cabify/pkg/entities"
)

func initStock(database *storage.Storage) (voucher, mug, shirt *entities.Product) {
	productStore := database.Products

	// Add Voucher
	voucher, _ = entities.NewProduct("", "VOUCHER", "Cabify Voucher", 500)
	productStore.Save(voucher)

	// Add Mug
	mug, _ = entities.NewProduct("", "MUG", "Cabify Mug", 750)
	productStore.Save(mug)

	// Add Shirt
	shirt, _ = entities.NewProduct("", "SHIRT", "Cabify Shirt", 2000)
	productStore.Save(shirt)

	return
}

func initPromotions(voucher, mug, shirt *entities.Product) []promotions.Promotion {
	return []promotions.Promotion{
		promotions.PayTwoGetOneFree(voucher.GetId()),
		promotions.BulkPurchase(shirt.GetId(), 1900),
	}
}

func main() {
	database, _ := internal.CreateStorage("memory")

	voucher, mug, shirt := initStock(database)
	initialPromotions := initPromotions(voucher, mug, shirt)

	mktplc := marketplace.NewMarketplace(database)
	checkout := marketplace.NewCheckout(database, initialPromotions)

	// Job Queue
	jobQueue := make(chan handlers.Job)
	go handlers.StartJobQueue(jobQueue, checkout)

	// Http server
	router := mux.NewRouter()

	var handler http.HandlerFunc

	handler = handlers.ListProducts(mktplc)
	router.HandleFunc("/products", handler).Methods("GET")

	handler = handlers.CreateBasket(jobQueue)
	router.HandleFunc("/checkouts", handler).Methods("POST")

	handler = handlers.DetailBasket(jobQueue)
	router.HandleFunc("/checkouts/{checkoutId}", handler).Methods("GET")

	handler = handlers.DeleteBasket(jobQueue)
	router.HandleFunc("/checkouts/{checkoutId}", handler).Methods("DELETE")

	handler = handlers.AddProduct(jobQueue)
	router.HandleFunc("/checkouts/{checkoutId}/products", handler).Methods("POST")

	handler = handlers.RemoveProduct(jobQueue)
	router.HandleFunc("/checkouts/{checkoutId}/products/{productId}", handler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

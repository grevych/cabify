package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/grevych/cabify/internal/handlers"
	mktplc "github.com/grevych/cabify/internal/marketplace"
	"github.com/grevych/cabify/internal/storage/memory"
	"github.com/grevych/cabify/pkg/entities"
)

func initStock(productStore *memory.ProductStore) {
	// Add Voucher
	voucher, _ := entities.NewProduct("", "VOUCHER", "Cabify Voucher", 400)
	productStore.Save(voucher)

	// Add Mug
	mug, _ := entities.NewProduct("", "MUG", "Cabify Mug", 500)
	productStore.Save(mug)

	// Add Shirt
	shirt, _ := entities.NewProduct("", "SHIRT", "Cabify Shirt", 600)
	productStore.Save(shirt)
}

func main() {
	jobQueue := make(chan handlers.Job)

	basketStore := memory.NewBasketStore()
	productStore := memory.NewProductStore()
	storage, _ := memory.NewStorage(basketStore, productStore)
	initStock(productStore)

	marketplace := mktplc.NewMarketplace(storage)
	checkout := mktplc.NewCheckout(storage)

	go handlers.StartJobQueue(jobQueue, checkout)

	router := mux.NewRouter()

	var handler http.HandlerFunc

	handler = handlers.ListProducts(marketplace)
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

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/grevych/cabify/internal/handlers"
	mktplc "github.com/grevych/cabify/internal/marketplace"
	stg "github.com/grevych/cabify/internal/storage"
)

func main() {
	jobQueue := make(chan string)
	storage := stg.Create("memory")

	marketplace := mktplc.NewMarketplace(storage)
	checkout := mktplc.NewCheckout(storage)

	go handlers.StartJobQueue(jobQueue)

	router := mux.NewRouter()

	handler := handlers.ListProducts(marketplace)
	router.HandleFunc("/products", handler).Methods("GET")

	handler := handlers.CreateBasket(checkout)
	router.HandleFunc("/checkouts", handler).Methods("POST")

	handler := handlers.DetailBasket(checkout, jobQueue)
	router.HandleFunc("/checkouts/{id}", handler).Methods("GET")

	handler := handlers.DeleteBasket(checkout, jobQueue)
	router.HandleFunc("/checkouts/{id}", handler).Methods("DELETE")

	handler := handlers.AddProduct(checkout, jobQueue)
	router.HandleFunc("/checkouts/{id}/products", handler).Methods("POST")

	handler := handlers.RemoveProduct(checkout, jobQueue)
	router.HandleFunc("/checkouts/{checkoutId}/products/{productId}", handler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}



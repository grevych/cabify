package handlers

import (
	"net/http"

	// "github.com/gorilla/mux"
	mktplc "github.com/grevych/cabify/internal/marketplace"
	"github.com/grevych/cabify/pkg/entities"
)

// vars := mux.Vars(r)
// title := vars["title"]
// page := vars["page"]

// fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

type job struct {
	action     string
	checkoutId string
	product    *entities.Product
}

type jobResponse struct {
}

func StartJobQueue(jobQueue chan job) {
	for j := range jobQueue {

	}
}

func ListProducts(marketplace *mktplc.Marketplace) http.HandlerFunc {
	// Doesn't need channel
	return func(w http.ResponseWriter, r *http.Request) {
		products := marketplace.ListProducts()
		r.Send()
	}
}

func CreateBasket(checkout *mktplc.Checkout) http.HandlerFunc {
	// Doesn't need channel
	return func(w http.ResponseWriter, r *http.Request) {
		basket := checkout.Create()
		r.Send()
	}
}

func DetailBasket(checkout *mktplc.Checkout, jobQueue chan job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		basket := checkout.Detail()
	}
}

func DeleteBasket(checkout *mktplc.Checkout, jobQueue chan job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		checkout.Delete()
	}
}

func AddProduct(checkout *mktplc.Checkout, jobQueue chan job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		checkout.AddProduct()
		r.Send()
	}
}

func RemoveProduct(checkout *mktplc.Checkout, jobQueue chan job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		checkout.RemoveProduct()
		r.Send()
	}
}

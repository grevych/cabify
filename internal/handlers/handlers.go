package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	mktplc "github.com/grevych/cabify/internal/marketplace"
	"github.com/grevych/cabify/pkg/entities"
)

type Job struct {
	action        string
	basketId      string
	productId     string
	response      *JobResponse
	responseQueue chan JobResponse
}

type JobResponse struct {
	err    error
	basket *entities.Basket
}

func StartJobQueue(jobQueue chan Job, checkout *mktplc.Checkout) {
	for job := range jobQueue {
		response := JobResponse{}

		if job.action == "create" {
			basket, err := checkout.Create()
			if err != nil {
				response.err = err
			} else {
				response.basket = basket
			}

		} else if job.action == "detail" {
			basket, err := checkout.Detail(job.basketId)
			if err != nil {
				response.err = err
			} else {
				response.basket = basket
			}

		} else if job.action == "delete" {
			basket, err := checkout.Detail(job.basketId)
			response.basket = basket
			if err != nil {
				response.err = err
			} else {
				response.err = checkout.Delete(job.basketId)
			}

		} else if job.action == "add_product" {
			err := checkout.AddProduct(job.basketId, job.productId)
			if err != nil {
				response.err = err
			} else {
				basket, _ := checkout.Detail(job.basketId)
				response.basket = basket
			}

		} else if job.action == "remove_product" {
			err := checkout.RemoveProduct(job.basketId, job.productId)
			if err != nil {
				response.err = err
			} else {
				basket, _ := checkout.Detail(job.basketId)
				response.basket = basket
			}
		}

		job.responseQueue <- response
	}
}

func ListProducts(marketplace *mktplc.Marketplace) http.HandlerFunc {
	// Doesn't need channel
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := marketplace.ListProducts()

		response, err := json.Marshal(products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func jobHandler(
	job *Job,
	jobQueue chan Job,
	jobResponseQueue chan JobResponse,
	w http.ResponseWriter,
) {
	jobQueue <- *job
	jobResponse := <-jobResponseQueue

	if jobResponse.err != nil {
		http.Error(w, jobResponse.err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(jobResponse.basket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateBasket(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		jobResponseQueue := make(chan JobResponse, 1)

		job := &Job{
			action:        "create",
			responseQueue: jobResponseQueue,
		}

		jobHandler(job, jobQueue, jobResponseQueue, w)
	}
}

func DetailBasket(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		jobResponseQueue := make(chan JobResponse, 1)
		vars := mux.Vars(r)

		job := &Job{
			action:        "detail",
			responseQueue: jobResponseQueue,
			basketId:      vars["checkoutId"],
		}

		jobHandler(job, jobQueue, jobResponseQueue, w)
	}
}

func DeleteBasket(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		jobResponseQueue := make(chan JobResponse, 1)
		vars := mux.Vars(r)

		job := &Job{
			action:        "delete",
			responseQueue: jobResponseQueue,
			basketId:      vars["checkoutId"],
		}

		jobHandler(job, jobQueue, jobResponseQueue, w)
	}
}

func AddProduct(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		jobResponseQueue := make(chan JobResponse, 1)
		vars := mux.Vars(r)

		var product struct{ Id string }
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &product)

		job := &Job{
			action:        "add_product",
			responseQueue: jobResponseQueue,
			basketId:      vars["checkoutId"],
			productId:     product.Id,
		}

		jobHandler(job, jobQueue, jobResponseQueue, w)
	}
}

func RemoveProduct(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		jobResponseQueue := make(chan JobResponse, 1)
		vars := mux.Vars(r)

		job := &Job{
			action:        "remove_product",
			responseQueue: jobResponseQueue,
			basketId:      vars["checkoutId"],
			productId:     vars["productId"],
		}

		jobHandler(job, jobQueue, jobResponseQueue, w)
	}
}

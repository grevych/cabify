package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grevych/cabify/internal/marketplace"
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

func StartJobQueue(jobQueue chan Job, checkout *marketplace.Checkout) {
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

func ListProducts(mktplc *marketplace.Marketplace) http.HandlerFunc {
	// Doesn't need channel
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		products, _ := mktplc.ListProducts()

		response, err := json.Marshal(products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	}
}

func CreateBasket(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jobResponseQueue := make(chan JobResponse, 1)

		job := &Job{
			action:        "create",
			responseQueue: jobResponseQueue,
		}

		jobQueue <- *job
		jobResponse := <-jobResponseQueue

		statusCode := manageErrorCode(jobResponse.err, 201)
		response, err := createResponse(jobResponse.basket, jobResponse.err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(statusCode)
		w.Write(response)
	}
}

func DetailBasket(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jobResponseQueue := make(chan JobResponse, 1)
		vars := mux.Vars(r)

		job := &Job{
			action:        "detail",
			responseQueue: jobResponseQueue,
			basketId:      vars["checkoutId"],
		}

		jobQueue <- *job
		jobResponse := <-jobResponseQueue

		statusCode := manageErrorCode(jobResponse.err, 200)
		response, err := createResponse(jobResponse.basket, jobResponse.err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(statusCode)
		w.Write(response)
	}
}

func DeleteBasket(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jobResponseQueue := make(chan JobResponse, 1)
		vars := mux.Vars(r)

		job := &Job{
			action:        "delete",
			responseQueue: jobResponseQueue,
			basketId:      vars["checkoutId"],
		}

		jobQueue <- *job
		jobResponse := <-jobResponseQueue

		statusCode := manageErrorCode(jobResponse.err, 200)
		response, err := createResponse(jobResponse.basket, jobResponse.err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(statusCode)
		w.Write(response)
	}
}

func AddProduct(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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

		jobQueue <- *job
		jobResponse := <-jobResponseQueue

		statusCode := manageErrorCode(jobResponse.err, 200)
		response, err := createResponse(jobResponse.basket, jobResponse.err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(statusCode)
		w.Write(response)
	}
}

func RemoveProduct(jobQueue chan Job) http.HandlerFunc {
	// Channel
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jobResponseQueue := make(chan JobResponse, 1)
		vars := mux.Vars(r)

		job := &Job{
			action:        "remove_product",
			responseQueue: jobResponseQueue,
			basketId:      vars["checkoutId"],
			productId:     vars["productId"],
		}

		jobQueue <- *job
		jobResponse := <-jobResponseQueue

		statusCode := manageErrorCode(jobResponse.err, 200)
		response, err := createResponse(jobResponse.basket, jobResponse.err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(statusCode)
		w.Write(response)
	}
}

func createResponse(basket *entities.Basket, err error) ([]byte, error) {
	if err != nil {
		errorResponse := struct {
			Message string
		}{
			Message: err.Error(),
		}

		return json.Marshal(&errorResponse)
	}

	return json.Marshal(basket)
}

func manageErrorCode(err error, expectedCode int) int {
	if err == nil {
		return expectedCode
	}

	if _, ok := err.(*marketplace.NotFoundError); ok {
		return http.StatusNotFound
	}

	if _, ok := err.(*marketplace.NotCreatedError); ok {
		return http.StatusConflict
	}

	if _, ok := err.(*marketplace.NotUpdatedError); ok {
		return http.StatusConflict
	}

	if _, ok := err.(*marketplace.NotDeletedError); ok {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}

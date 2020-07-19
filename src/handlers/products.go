package handlers

import (
	"context"
	"log"
	"net/http"
	"products-api/src/data"
	"strconv"

	"github.com/gorilla/mux"
)

// Products handler
type Products struct {
	logger *log.Logger
}

// NewProducts creates a new products handler
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

// KeyProduct is the context key for where the extracted product is stored
type KeyProduct struct{}

// ProductValidationMiddleware extracts the product from the body
func ProductValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := data.Product{}
		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		nextRequest := r.WithContext(ctx)
		next.ServeHTTP(w, nextRequest)
	})
}

// GetProducts fetches all the products from the datasource
func (products *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	productsCollections := data.GetProducts()
	err := productsCollections.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct upload a new products to the datasource
func (products *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	products.logger.Println("post a new product")
	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)
	return
}

// UpdateProduct re-writes a product by id
func (products *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	products.logger.Println("handle PUT request")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "unable to parse id", http.StatusBadRequest)
		return
	}

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	} else if err != nil {
		// the product was found but something else went wrong
		http.Error(w, "could not update the product", http.StatusInternalServerError)
		return
	}

	return
}

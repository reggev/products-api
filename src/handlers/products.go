package handlers

import (
	"log"
	"net/http"
	"products-api/src/data"
)

// Products handler
type Products struct {
	logger *log.Logger
}

// NewProducts creates a new products handler
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	productsCollections := data.GetProducts()
	err := productsCollections.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (products *Products) addProduce(w http.ResponseWriter, r *http.Request) {
	products.logger.Println("post a new product")
	return
}

func (products *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		products.getProducts(w, r)
		return
	} else if r.Method == http.MethodPost {
		products.addProduce(w, r)
		return
	}
	// TODO:: handle an update
	w.WriteHeader(http.StatusMethodNotAllowed)
}

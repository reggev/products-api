package handlers

import (
	"log"
	"net/http"
	"products-api/src/data"
	"regexp"
	"strconv"
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
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	products.logger.Printf("\nProduct: %#v", product)
	data.AddProduct(product)
	return
}

func (products *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	products.logger.Println("handle PUT request")
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	products.logger.Printf("\nProduct: %#v", product)
	err = data.UpdateProduct(id, product)
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

func (products *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		products.getProducts(w, r)
		return
	} else if r.Method == http.MethodPost {
		products.addProduce(w, r)
		return
	} else if r.Method == http.MethodPut {
		// expect the id in the uri
		// regexp.MustCompile validates the regex on initialization - this is a safety method
		regex := regexp.MustCompile(`/([0-9]+)`)
		// the second argument is number of results at most for any n>=0, -1 disables the limit
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)
		products.logger.Println(group)
		if len(group) != 1 || len(group[0]) != 2 {
			products.logger.Println("invalid group or capture length ::", group)
			http.Error(w, "invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			products.logger.Println("could not convert the id to a number ::", idString)
			http.Error(w, "could not convert the id to a number", http.StatusBadRequest)
		}
		products.updateProduct(id, w, r)
		return
	}
	// TODO:: handle an update
	w.WriteHeader(http.StatusMethodNotAllowed)
}

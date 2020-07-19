package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product properties
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products represents a slice of Product
type Products []*Product

// ToJSON encodes the products to json
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// Validate json values after parsing
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)
	return validate.Struct(p)
}

func validateSku(fieldLevel validator.FieldLevel) bool {
	// sku format abc-abcd-abcdef
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	matches := re.FindAllString(fieldLevel.Field().String(), -1)
	return len(matches) == 1
}

// FromJSON decodes a JSON request to a Products type
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

// GetProducts fetch a list of products
func GetProducts() Products {
	return productList
}

// ErrProductNotFound id an error in case of findProduct fail
var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for idx, product := range productList {
		if product.ID == id {
			return product, idx, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// UpdateProduct finds and updates an existing product
func UpdateProduct(id int, product *Product) error {
	currentProduct, position, err := findProduct(id)
	if err != nil {
		return err
	}
	product.ID = currentProduct.ID
	productList[position] = product
	return nil
}

// AddProduct adds a product to the products collection
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd43",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

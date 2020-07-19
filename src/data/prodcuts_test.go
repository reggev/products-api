package data

import (
	"testing"
)

// TestChecksValidation validates a product
func TestChecksValidation(t *testing.T) {
	product := &Product{
		Name:  "joe",
		Price: 2,
		SKU:   "abc-def-ghi",
	}
	err := product.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

// Package `core` provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import "github.com/dominikbraun/foodunit/dl"

// `SupplierInteractor` represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type SupplierInteractor struct {
	r dl.SupplierRepository
}

// Creates a new `SupplierInteractor` instance.
func NewSupplierInteractor(r dl.SupplierRepository) *SupplierInteractor {
	return &SupplierInteractor{r: r}
}

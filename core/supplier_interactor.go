// Package `core` provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import (
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/dl"
)

// `SupplierInteractor` represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type SupplierInteractor struct {
	suppliers  dl.SupplierRepository
	categories dl.CategoryRepository
	dishes     dl.DishRepository
}

// Creates a new `SupplierInteractor` instance.
func NewSupplierInteractor() *SupplierInteractor {
	return &SupplierInteractor{
		suppliers:  load.SupplierRepository(),
		categories: load.CategoryRepository(),
		dishes:     load.DishRepository(),
	}
}

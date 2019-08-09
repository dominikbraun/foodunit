// Package `core` provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import (
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/pkg/errors"
)

// `SupplierInteractor` represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type SupplierInteractor struct {
	suppliers  dl.SupplierRepository
	categories dl.CategoryRepository
	dishes     dl.DishRepository
}

// Creates a new `SupplierInteractor` instance.
func NewSupplierInteractor() (*SupplierInteractor, error) {
	loader := load.RepositoryLoader

	if !loader.IsReady {
		return nil, errors.New("repository loader has to be initialized first")
	}

	return &SupplierInteractor{
		suppliers:  loader.Suppliers(),
		categories: loader.Categories(),
		dishes:     loader.Dishes(),
	}, nil
}

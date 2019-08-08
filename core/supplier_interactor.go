// Package `core` provides types for interacting with the core
// business logic. It depicts the default use case of FoodUnit.
package core

import "github.com/dominikbraun/foodunit/dl"

type SupplierInteractor struct {
	r dl.SupplierRepository
}

func NewSupplierInteractor(r dl.SupplierRepository) *SupplierInteractor {
	return &SupplierInteractor{r: r}
}

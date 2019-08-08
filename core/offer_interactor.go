// Package `core` provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import "github.com/dominikbraun/foodunit/dl"

// `OfferInteractor` represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type OfferInteractor struct {
	r dl.OfferRepository
}

// Creates a new `OfferInteractor` instance.
func NewOfferInteractor(r dl.OfferRepository) *OfferInteractor {
	return &OfferInteractor{r: r}
}

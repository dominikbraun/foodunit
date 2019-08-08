// Package `core` provides types for interacting with the core
// business logic. It depicts the default use case of FoodUnit.
package core

import "github.com/dominikbraun/foodunit/dl"

type OfferInteractor struct {
	r dl.OfferRepository
}

func NewOfferInteractor(r dl.OfferRepository) *OfferInteractor {
	return &OfferInteractor{r: r}
}

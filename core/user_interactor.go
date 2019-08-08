// Package `core` provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import "github.com/dominikbraun/foodunit/dl"

// `UserInteractor` represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type UserInteractor struct {
	r dl.UserRepository
}

// Creates a new `UserInteractor` instance.
func NewUserInteractor(r dl.UserRepository) *UserInteractor {
	return &UserInteractor{r: r}
}

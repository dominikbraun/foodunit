// Package `core` provides types for interacting with the core
// business logic. It depicts the default use case of FoodUnit.
package core

import "github.com/dominikbraun/foodunit/dl"

type UserInteractor struct {
	r dl.UserRepository
}

func NewUserInteractor(r dl.UserRepository) *UserInteractor {
	return &UserInteractor{r: r}
}

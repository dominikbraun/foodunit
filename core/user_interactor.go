// Package core provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import (
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/pkg/errors"
)

// UserInteractor represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type UserInteractor struct {
	users dl.UserRepository
}

// Creates a new UserInteractor instance.
func NewUserInteractor(r dl.UserRepository) (*UserInteractor, error) {
	loader := load.RepositoryLoader

	if !loader.IsReady {
		return nil, errors.New("repository loader has to be initialized first")
	}

	return &UserInteractor{
		users: loader.Users,
	}, nil
}

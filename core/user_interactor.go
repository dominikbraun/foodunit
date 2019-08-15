// Copyright 2019 The FoodUnit Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package core provides types for interacting with the core business logic.
// It depicts the default use case of FoodUnit.
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

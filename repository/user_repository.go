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

// Package repository provides discrete DL repository implementations.
package repository

import (
	"github.com/dominikbraun/foodunit/dl"
	"github.com/dominikbraun/foodunit/storage"
)

// UserRepository represents the default repository for managing Users.
type UserRepository struct {
	storage storage.Repository
}

// Create implements dl.UserRepository.Create.
func (u UserRepository) Create(user dl.User) error {
	err := u.storage.Create(user)
	return err
}

// Find implements dl.UserRepository.Find.
func (u UserRepository) Find(id uint64) (dl.User, error) {
	var user dl.User
	err := u.storage.Find(id, &user)

	return user, err
}

// Update implements dl.UserRepository.Update.
func (u UserRepository) Update(user dl.User) error {
	err := u.storage.Update(user)
	return err
}

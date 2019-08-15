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

// Package mariadb provides repository implementations as SQL gateways.
package mariadb

import (
	"github.com/dominikbraun/foodunit/dl"
)

type (
	// Repository holds methods of dl.Repository.
	UserRepository struct{}
)

// Migrate implements dl.UserRepository.Migrate.
func (u UserRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.UserRepository.Create.
func (u UserRepository) Create(user *dl.User) (uint64, error) {
	panic("implement me")
}

// Find implements dl.UserRepository.Find.
func (u UserRepository) Find(id uint) (*dl.User, error) {
	panic("implement me")
}

// FindByMail implements dl.UserRepository.FindByMail.
func (u UserRepository) FindByMail(mail string) (*dl.User, error) {
	panic("implement me")
}

// Update implements dl.UserRepository.Update.
func (u UserRepository) Update(user *dl.User) error {
	panic("implement me")
}

// Delete implements dl.UserRepository.Delete.
func (u UserRepository) Delete(user *dl.User) error {
	panic("implement me")
}

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

// Package storage provides storage implementations for managing entities.
package storage

import (
	"github.com/dominikbraun/foodunit/dl"
)

// Repository provides methods to perform CRUD operations on entities.
// Mapping the entity type against a table or bucket is up to the implementation.
type Repository interface {
	Create(dl.Entity) error
	Find(id uint64, dest dl.Entity) error
	FindBy(field, val string, dests *[]dl.Entity) error
	Update(dl.Entity) error
	Delete(dl.Entity) error
}
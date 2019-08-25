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

// Manager represents a wrapper for all available DL repository implementations.
// It can be used to inject a storage object once for all repositories.
type Manager struct {
	suppliers       dl.SupplierRepository
	dishes          dl.DishRepository
	characteristics dl.CharacteristicRepository
}

// NewManager creates and returns a Manager instance.
func NewManager(storage storage.Repository) Manager {
	manager := Manager{
		suppliers:       SupplierRepository{storage: storage},
		dishes:          DishRepository{storage: storage},
		characteristics: CharacteristicRepository{storage: storage},
	}
	return manager
}

// Suppliers returns the default SupplierRepository implementation.
func (m Manager) Suppliers() dl.SupplierRepository {
	return m.suppliers
}

// Dishes returns the default DishesRepository implementation.
func (m Manager) Dishes() dl.DishRepository {
	return m.dishes
}

// Characteristics returns the default CharacteristicsRepository implementation.
func (m Manager) Characteristics() dl.CharacteristicRepository {
	return m.characteristics
}

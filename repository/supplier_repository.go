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
	"github.com/pkg/errors"
)

// SupplierRepository represents the default repository for managing Suppliers.
type SupplierRepository struct {
	storage storage.Repository
}

// Create implements dl.SupplierRepository.Create.
func (s SupplierRepository) Create(supplier dl.Supplier) error {
	err := s.storage.Create(supplier)
	return err
}

// Find implements dl.SupplierRepository.Find.
func (s SupplierRepository) Find(id uint64) (dl.Supplier, error) {
	var supplier dl.Supplier
	err := s.storage.Find(id, &supplier)

	return supplier, err
}

// FindBy implements dl.SupplierRepository.FindBy.
func (s SupplierRepository) FindBy(field, val string) ([]dl.Supplier, error) {
	var entities []dl.Entity

	if err := s.storage.FindBy(field, val, &entities); err != nil {
		return nil, err
	}

	suppliers := make([]dl.Supplier, len(entities))

	for i, e := range entities {
		supplier, ok := e.(dl.Supplier)
		if !ok {
			return nil, errors.New("received entity is not of type Supplier")
		}
		suppliers[i] = supplier
	}
	return suppliers, nil
}

// Update implements dl.SupplierRepository.Update.
func (s SupplierRepository) Update(supplier dl.Supplier) error {
	err := s.storage.Update(supplier)
	return err
}

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

// Package sql provides a dl.DataAccess implementation for MariaDB.
package sql

import "github.com/dominikbraun/foodunit/dl"

type (
	// supplierRepository implements dl.SupplierRepository.
	supplierRepository struct{}
	// categoryRepository implements dl.CategoryRepository.
	categoryRepository struct{}
	// dishRepository implements dl.DishRepository.
	dishRepository struct{}
	// characteristicRepository implements dl.CharacteristicRepository.
	characteristicRepository struct{}
	// variantRepository implements dl.VariantRepository.
	variantRepository struct{}
)

// Migrate implements dl.SupplierRepository.Migrate.
func (s supplierRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.SupplierRepository.Create.
func (s supplierRepository) Create(supplier *dl.Supplier) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.SupplierRepository.FindByID.
func (s supplierRepository) FindByID(id uint64) (dl.Supplier, error) {
	panic("implement me")
}

// Update implements dl.SupplierRepository.Update.
func (s supplierRepository) Update(supplier *dl.Supplier) error {
	panic("implement me")
}

// Delete implements dl.SupplierRepository.Delete.
func (s supplierRepository) Delete(supplier *dl.Supplier) error {
	panic("implement me")
}

// Migrate implements dl.CategoryRepository.Migrate.
func (c categoryRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.CategoryRepository.Create.
func (c categoryRepository) Create(category *dl.Category) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.CategoryRepository.FindByID.
func (c categoryRepository) FindByID(id uint64) (dl.Category, error) {
	panic("implement me")
}

// FindBySupplierID implements dl.CategoryRepository.FindBySupplierID.
func (c categoryRepository) FindBySupplierID(supplierID uint64) ([]dl.Category, error) {
	panic("implement me")
}

// Update implements dl.CategoryRepository.Update.
func (c categoryRepository) Update(category *dl.Category) error {
	panic("implement me")
}

// Delete implements dl.CategoryRepository.Delete.
func (c categoryRepository) Delete(category *dl.Category) error {
	panic("implement me")
}

// Migrate implements dl.DishRepository.Migrate.
func (d dishRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.DishRepository.Create.
func (d dishRepository) Create(dish *dl.Dish) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.DishRepository.FindByID.
func (d dishRepository) FindByID(id uint64) (dl.Dish, error) {
	panic("implement me")
}

// FindByCategoryID implements dl.DishRepository.FindByCategoryID.
func (d dishRepository) FindByCategoryID(categoryID uint64) ([]dl.Dish, error) {
	panic("implement me")
}

// Update implements dl.DishRepository.Update.
func (d dishRepository) Update(dish *dl.Dish) error {
	panic("implement me")
}

// Delete implements dl.DishRepository.Delete.
func (d dishRepository) Delete(dish *dl.Dish) error {
	panic("implement me")
}

// Migrate implements dl.CharacteristicRepository.Migrate.
func (c characteristicRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.CharacteristicRepository.Create.
func (c characteristicRepository) Create(characteristic *dl.Characteristic) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.CharacteristicRepository.FindByID.
func (c characteristicRepository) FindByID(id uint64) (dl.Characteristic, error) {
	panic("implement me")
}

// Update implements dl.CharacteristicRepository.Update.
func (c characteristicRepository) Update(characteristic *dl.Characteristic) error {
	panic("implement me")
}

// Delete implements dl.CharacteristicRepository.Delete.
func (c characteristicRepository) Delete(characteristic *dl.Characteristic) error {
	panic("implement me")
}

// Migrate implements dl.VariantRepository.Migrate.
func (v variantRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.VariantRepository.Create.
func (v variantRepository) Create(variant *dl.Variant) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.VariantRepository.FindByID.
func (v variantRepository) FindByID(id uint64) (dl.Variant, error) {
	panic("implement me")
}

// FindByCharacteristicID implements dl.VariantRepository.FindByCharacteristicID.
func (v variantRepository) FindByCharacteristicID(id uint64) ([]dl.Variant, error) {
	panic("implement me")
}

// Update implements dl.VariantRepository.Update.
func (v variantRepository) Update(variant *dl.Variant) error {
	panic("implement me")
}

// Delete implements dl.VariantRepository.Delete.
func (v variantRepository) Delete(variant *dl.Variant) error {
	panic("implement me")
}

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

// DishRepository represents the default repository for managing Dishes.
type DishRepository struct {
	storage storage.Repository
}

// Create implements dl.DishRepository.Create.
func (d DishRepository) Create(dish dl.Dish) error {
	err := d.storage.Create(dish)
	return err
}

// FindBy implements dl.DishRepository.FindBy.
func (d DishRepository) FindBy(field, val string) ([]dl.Dish, error) {
	var entities []dl.Entity

	if err := d.storage.FindBy(field, val, &entities); err != nil {
		return nil, err
	}

	suppliers := make([]dl.Dish, len(entities))

	for i, e := range entities {
		dish, ok := e.(dl.Dish)
		if !ok {
			return nil, errors.New("received entity is not of type Dish")
		}
		suppliers[i] = dish
	}
	return suppliers, nil
}

// Update implements dl.DishRepository.Update.
func (d DishRepository) Update(dish dl.Dish) error {
	err := d.storage.Update(dish)
	return err
}

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

// CharacteristicRepository represents the default repository for managing Characteristics.
type CharacteristicRepository struct {
	storage storage.Repository
}

// Create implements dl.DishRepository.Create.
func (c CharacteristicRepository) Create(characteristic dl.Characteristic) error {
	err := c.storage.Create(characteristic)
	return err
}

// FindBy implements dl.DishRepository.FindBy.
func (c CharacteristicRepository) FindBy(field, val string) ([]dl.Characteristic, error) {
	var entities []dl.Entity

	if err := c.storage.FindBy(field, val, &entities); err != nil {
		return nil, err
	}

	characteristics := make([]dl.Characteristic, len(entities))

	for i, e := range entities {
		characteristic, ok := e.(dl.Characteristic)
		if !ok {
			return nil, errors.New("received entity is not of type Characteristic")
		}
		characteristics[i] = characteristic
	}
	return characteristics, nil
}

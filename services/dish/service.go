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

// Package dish provides services and types for Dish-related data.
package dish

import (
	"database/sql"
	"errors"
	"github.com/dominikbraun/foodunit/storage"
)

var (
	ErrCharacteristicsNotFound = errors.New("the characteristics for the dish could not be found")
	ErrVariantsNotFound        = errors.New("the variants for the characteristics could not be found")
)

// Service executes dish-related business logic and use cases. It is also responsible
// for accessing the model storage under consideration of all business rules.
type Service struct {
	dishes          storage.Dish
	characteristics storage.Characteristic
	variants        storage.Variant
}

// NewService creates a new Service instance utilizing the given storage objects.
// The storage objects need to be ready to use for the service.
func NewService(d storage.Dish, c storage.Characteristic, v storage.Variant) *Service {
	service := Service{
		dishes:          d,
		characteristics: c,
		variants:        v,
	}
	return &service
}

// GetCharacteristics returns all characteristics of a given dish.
func (s *Service) GetCharacteristics(dishID uint64) ([]Characteristic, error) {
	characteristicEntities, err := s.characteristics.FindByDish(dishID)

	if err == sql.ErrNoRows {
		return nil, ErrCharacteristicsNotFound
	} else if err != nil {
		return nil, err
	}

	characteristics := make([]Characteristic, 0)

	// ToDo: Load the variants via SQL in the corresponding storage
	for _, characteristicEntity := range characteristicEntities {
		variantEntities, err := s.variants.FindByCharacteristic(characteristicEntity.ID)

		if err == sql.ErrNoRows {
			return nil, ErrVariantsNotFound
		} else if err != nil {
			return nil, err
		}

		variants := make([]Variant, 0)

		for _, variantEntity := range variantEntities {
			variants = append(variants, Variant{
				ID:        variantEntity.ID,
				Value:     variantEntity.Value,
				IsDefault: variantEntity.IsDefault,
				Price:     variantEntity.Price,
			})
		}

		characteristics = append(characteristics, Characteristic{
			ID:       characteristicEntity.ID,
			Name:     characteristicEntity.Name,
			Multiple: characteristicEntity.Multiple,
			Variants: variants,
		})
	}

	return characteristics, nil
}

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

// Package core provides business services and use case methods.
package core

import (
	"github.com/dominikbraun/foodunit/repository"
	"strconv"
)

// SupplierInteractor exposes methods to execute Supplier-related use cases.
type SupplierInteractor struct {
	repository repository.Manager
}

// GetInfo implements dl.SupplierService.GetInfo.
func (s SupplierInteractor) GetInfo(id uint64) (Supplier, error) {
	rs, err := s.repository.Suppliers().Find(id)
	if err != nil {
		return Supplier{}, nil
	}

	supplier := Supplier{rs}
	return supplier, nil
}

// GetMenu implements dl.SupplierService.GetMenu.
func (s SupplierInteractor) GetMenu(id uint64) (Menu, error) {
	supplierID := strconv.Itoa(int(id))

	rcs, err := s.repository.Dishes().FindCategoryBy("supplier_id", supplierID)
	if err != nil {
		return nil, err
	}

	var menu Menu

	for _, c := range rcs {
		var dishes []Dish

		categoryID := strconv.Itoa(int(c.ID))

		rds, err := s.repository.Dishes().FindBy("category_id", categoryID)
		if err != nil {
			return nil, err
		}

		for _, d := range rds {
			dish := Dish{
				Name:        d.Name,
				Description: d.Description,
				Price:       d.Price,
				IsUncertain: d.IsUncertain,
			}
			dishes = append(dishes, dish)
		}

		category := Category{
			Name:    c.Name,
			ImgPath: c.ImgPath,
			Dishes:  dishes,
		}

		menu = append(menu, category)
	}
	return menu, nil
}

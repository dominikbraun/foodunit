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

// Package core provides business logic interactors and services.
package core

import (
	"github.com/dominikbraun/foodunit/core/dto"
	"github.com/dominikbraun/foodunit/storage"
	"time"
)

// GetRestaurantInfo is the public interface for retrieving meta information
// for a given restaurant, returning an instance of core.RestaurantInfo.
func GetRestaurantInfo(id uint64, model storage.RestaurantModel) (dto.RestaurantInfo, error) {
	r, err := model.GetInfo(id)
	if err != nil {
		return dto.RestaurantInfo{}, err
	}

	var open string

	switch time.Now().Weekday() {
	case time.Monday:
		open = r.OpenMon
	case time.Tuesday:
		open = r.OpenTue
	case time.Wednesday:
		open = r.OpenWed
	case time.Thursday:
		open = r.OpenThu
	case time.Friday:
		open = r.OpenFri
	case time.Saturday:
		open = r.OpenSat
	case time.Sunday:
		open = r.OpenSun
	}

	ri := dto.RestaurantInfo{
		Name:       r.Name,
		Street:     r.Street,
		PostalCode: r.PostalCode,
		City:       r.City,
		Phone:      r.Phone,
		Open:       open,
		Website:    r.Website,
	}

	return ri, nil
}

// GetRestaurantMenu is the most convenient way to retrieve a restaurant's menu. It
// holds all menu sections (categories) and the corresponding dishes, including prices
// and meta information. It does not include the dish characteristic variants.
func GetRestaurantMenu(id uint64, categoryModel storage.CategoryModel, dishModel storage.DishModel) (dto.Menu, error) {
	categories, err := categoryModel.FindByRestaurant(id)
	if err != nil {
		return dto.Menu{}, err
	}

	var menuCategories []dto.MenuCategory

	for _, c := range categories {
		dishes, err := dishModel.FindByCategory(c.ID)
		if err != nil {
			return dto.Menu{}, err
		}

		var menuDishes []dto.MenuDish

		for _, d := range dishes {
			menuDish := dto.MenuDish{
				Name:         d.Name,
				Description:  d.Description,
				Price:        d.Price,
				IsUncertain:  d.IsUncertain,
				IsHealthy:    d.IsHealthy,
				IsVegetarian: d.IsVegetarian,
			}
			menuDishes = append(menuDishes, menuDish)
		}

		menuCategory := dto.MenuCategory{
			Name:   c.Name,
			Dishes: menuDishes,
		}
		menuCategories = append(menuCategories, menuCategory)
	}

	menu := dto.Menu{
		Categories: menuCategories,
	}

	return menu, nil
}

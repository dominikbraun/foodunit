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

// Package restaurant provides services and types for Restaurant-related data.
package restaurant

import (
	"database/sql"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrRestaurantNotFound = errors.New("the restaurant could not be found")
	ErrMenuNotFound       = errors.New("the menu could not be found")
)

type Service struct {
	restaurants storage.Restaurant
	categories  storage.Category
	dishes      storage.Dish
}

func NewService(r storage.Restaurant, c storage.Category, d storage.Dish) *Service {
	service := Service{
		restaurants: r,
		categories:  c,
		dishes:      d,
	}
	return &service
}

func (s *Service) Info(id uint64) (Info, error) {
	restaurant, err := s.restaurants.Find(id)

	if err == sql.ErrNoRows {
		return Info{}, ErrRestaurantNotFound
	} else if err != nil {
		return Info{}, err
	}

	var open string

	switch time.Now().Weekday() {
	case time.Monday:
		open = restaurant.OpenMon
	case time.Tuesday:
		open = restaurant.OpenTue
	case time.Wednesday:
		open = restaurant.OpenWed
	case time.Thursday:
		open = restaurant.OpenThu
	case time.Friday:
		open = restaurant.OpenFri
	case time.Saturday:
		open = restaurant.OpenSat
	case time.Sunday:
		open = restaurant.OpenSun
	}

	info := Info{
		Name:       restaurant.Name,
		Street:     restaurant.Street,
		PostalCode: restaurant.PostalCode,
		City:       restaurant.City,
		Phone:      restaurant.Phone,
		Open:       open,
		Website:    restaurant.Website,
	}

	return info, nil
}

func (s *Service) Menu(id uint64) (Menu, error) {
	categories, err := s.categories.FindByRestaurant(id)

	if err == sql.ErrNoRows {
		return Menu{}, ErrMenuNotFound
	} else if err != nil {
		return Menu{}, err
	}

	var menuCategories []MenuCategory

	for _, c := range categories {
		dishes, err := s.dishes.FindByCategory(c.ID)

		if err == sql.ErrNoRows {
			return Menu{}, ErrMenuNotFound
		} else if err != nil {
			return Menu{}, err
		}

		var menuDishes []MenuDish

		for _, d := range dishes {
			menuDish := MenuDish{
				Name:         d.Name,
				Description:  d.Description,
				Price:        d.Price,
				IsUncertain:  d.IsUncertain,
				IsHealthy:    d.IsHealthy,
				IsVegetarian: d.IsVegetarian,
			}
			menuDishes = append(menuDishes, menuDish)
		}

		menuCategory := MenuCategory{
			Name:   c.Name,
			Dishes: menuDishes,
		}
		menuCategories = append(menuCategories, menuCategory)
	}

	menu := Menu{
		Categories: menuCategories,
	}

	return menu, nil
}

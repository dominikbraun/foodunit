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

// Service executes restaurant-related business logic and use cases. It is also responsible
// for accessing the model storage under consideration of all business rules.
type Service struct {
	restaurants storage.Restaurant
	categories  storage.Category
	dishes      storage.Dish
}

// NewService creates a new Service instance utilizing the given storage objects.
// The storage objects need to be ready to use for the service.
func NewService(r storage.Restaurant, c storage.Category, d storage.Dish) *Service {
	service := Service{
		restaurants: r,
		categories:  c,
		dishes:      d,
	}
	return &service
}

// Info returns relevant meta data for an restaurant, e. g. the opening hours. It
// won't return the restaurant's menu (see Menu for this).
func (s *Service) Info(id uint64) (Info, error) {
	restaurantEntity, err := s.restaurants.Find(id)

	if err == sql.ErrNoRows {
		return Info{}, ErrRestaurantNotFound
	} else if err != nil {
		return Info{}, err
	}

	var open string

	switch time.Now().Weekday() {
	case time.Monday:
		open = restaurantEntity.OpenMon
	case time.Tuesday:
		open = restaurantEntity.OpenTue
	case time.Wednesday:
		open = restaurantEntity.OpenWed
	case time.Thursday:
		open = restaurantEntity.OpenThu
	case time.Friday:
		open = restaurantEntity.OpenFri
	case time.Saturday:
		open = restaurantEntity.OpenSat
	case time.Sunday:
		open = restaurantEntity.OpenSun
	}

	info := Info{
		Name:       restaurantEntity.Name,
		Street:     restaurantEntity.Street,
		PostalCode: restaurantEntity.PostalCode,
		City:       restaurantEntity.City,
		Phone:      restaurantEntity.Phone,
		Open:       open,
		Website:    restaurantEntity.Website,
	}

	return info, nil
}

// Menu returns a restaurant's menu including all categories or menu sections and
// their individual dishes. However, the dish characteristics are not included here
// due to performance reasons.
func (s *Service) Menu(id uint64) (Menu, error) {
	categoryEntities, err := s.categories.FindByRestaurant(id)

	if err == sql.ErrNoRows || len(categoryEntities) == 0 {
		return Menu{}, ErrMenuNotFound
	} else if err != nil {
		return Menu{}, err
	}

	menuCategories := make([]MenuCategory, 0)

	for _, c := range categoryEntities {
		dishEntities, err := s.dishes.FindByCategory(c.ID)

		if err == sql.ErrNoRows || len(dishEntities) == 0 {
			return Menu{}, ErrMenuNotFound
		} else if err != nil {
			return Menu{}, err
		}

		menuDishes := make([]MenuDish, 0)

		for _, d := range dishEntities {
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

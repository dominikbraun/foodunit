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
	"github.com/dominikbraun/foodunit/storage"
	"time"
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
	if err != nil {
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

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

// Package order provides services and types for Order-related data.
package order

import (
	"database/sql"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/pkg/errors"
)

var (
	ErrOfferNotFound = errors.New("the offer could not be found")
)

type Service struct {
	orders          storage.Order
	positions       storage.Position
	configurations  storage.Configuration
	dishes          storage.Dish
	characteristics storage.Characteristic
	variants        storage.Variant
}

func NewService(o storage.Order, p storage.Position, c storage.Configuration, d storage.Dish, chr storage.Characteristic, v storage.Variant) *Service {
	service := Service{
		orders:          o,
		positions:       p,
		configurations:  c,
		dishes:          d,
		characteristics: chr,
		variants:        v,
	}
	return &service
}

func (s *Service) GetAll(offerID uint64) ([]Order, error) {
	orders, err := s.orders.FindByOffer(offerID)

	if err == sql.ErrNoRows {
		return []Order{}, ErrOfferNotFound
	} else if err != nil {
		return []Order{}, err
	}

	allOrders := make([]Order, 0)

	for _, o := range orders {
		order := Order{
			ID:     o.ID,
			User:   User{Name: o.User.Name},
			IsPaid: o.IsPaid,
			Total:  0,
		}

		positions, err := s.positions.FindByOrder(o.ID)

		if err == sql.ErrNoRows {
			continue
		} else if err != nil {
			return []Order{}, err
		}

		for _, p := range positions {
			dish, err := s.dishes.Find(p.Dish.ID)
			alternative, err := s.dishes.Find(p.Alternative.ID)

			configurations, err := s.configurations.FindByPosition(p.ID)

			if err == sql.ErrNoRows {
				continue
			} else if err != nil {
				return []Order{}, err
			}

			positionConfigurations := make([]Configuration, 0)

			for _, c := range configurations {
				variants, err := s.configurations.FindVariants(c.ID)

				if err == sql.ErrNoRows {
					continue
				} else if err != nil {
					return []Order{}, err
				}

				configurationVariants := make([]Variant, 0)

				for _, v := range variants {
					variant := Variant{
						ID:    v.ID,
						Value: v.Value,
						Price: v.Price,
					}

					configurationVariants = append(configurationVariants, variant)
				}

				configuration := Configuration{
					CharacteristicName: c.Characteristic.Name,
					Multiple:           c.Characteristic.Multiple,
					Variants:           configurationVariants,
				}

				positionConfigurations = append(positionConfigurations, configuration)
			}

			position := Position{
				ID: p.ID,
				Dish: Dish{
					ID:    dish.ID,
					Name:  dish.Name,
					Price: dish.Price,
				},
				Alternative: Dish{
					ID:    alternative.ID,
					Name:  alternative.Name,
					Price: alternative.Price,
				},
				Note:           p.Note,
				Configurations: positionConfigurations,
			}

			order.Positions = append(order.Positions, position)
			order.Total += dish.Price
		}

		allOrders = append(allOrders, order)
	}

	return allOrders, nil
}

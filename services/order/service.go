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
	"github.com/dominikbraun/foodunit/model"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/pkg/errors"
)

var (
	ErrOfferNotFound = errors.New("the offer could not be found")
	ErrOrderNotFound = errors.New("the order could not be found")
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
		order, err := s.buildOrder(&o)
		if err != nil {
			return []Order{}, err
		}

		allOrders = append(allOrders, order)
	}

	return allOrders, nil
}

func (s *Service) Get(offerID, userID uint64) (Order, error) {
	orderEntry, err := s.orders.FindByOfferAndUser(offerID, userID)

	if err == sql.ErrNoRows {
		return Order{}, ErrOrderNotFound
	} else if err != nil {
		return Order{}, err
	}

	order, err := s.buildOrder(&orderEntry)

	return order, err
}

func (s *Service) buildOrder(orderEntry *model.Order) (Order, error) {
	order := Order{
		ID:     orderEntry.ID,
		User:   User{Name: orderEntry.User.Name},
		IsPaid: orderEntry.IsPaid,
		Total:  0,
	}

	positions, err := s.positions.FindByOrder(orderEntry.ID)

	if err == sql.ErrNoRows && err != nil {
		return Order{}, err
	}

	for _, p := range positions {
		dish, err := s.dishes.Find(p.Dish.ID)
		alternative, err := s.dishes.Find(p.Alternative.ID)
		configurations, err := s.loadConfigurations(p.ID)

		if err == sql.ErrNoRows {
			continue
		} else if err != nil {
			return Order{}, err
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
			Configurations: configurations,
		}

		order.Positions = append(order.Positions, position)
		order.Total += dish.Price
	}

	return order, nil
}

func (s *Service) loadConfigurations(positionID uint64) ([]Configuration, error) {
	configs, err := s.configurations.FindByPosition(positionID)

	if err == sql.ErrNoRows {
		return []Configuration{}, nil
	} else if err != nil {
		return []Configuration{}, err
	}

	configurations := make([]Configuration, 0)

	for _, c := range configs {
		variants, err := s.configurations.FindVariants(c.ID)

		if err == sql.ErrNoRows {
			continue
		} else if err != nil {
			return []Configuration{}, err
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

		configurations = append(configurations, configuration)
	}

	return configurations, nil
}

func (s *Service) Update(order *Update) error {
	_, err := s.orders.FindByOfferAndUser(order.OfferID, order.UserID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	orderEntry := model.Order{
		User:   model.User{ID: order.UserID},
		IsPaid: false,
	}
	orderID, err := s.orders.Store(order.OfferID, &orderEntry)
	if err != nil {
		return err
	}

	for _, p := range order.Positions {
		positionEntry := model.Position{
			Dish:        model.Dish{ID: p.DishID},
			Alternative: model.Dish{ID: p.AlternativeDishID},
			Note:        p.Note,
		}
		positionID, err := s.positions.Store(orderID, &positionEntry)
		if err != nil {
			return err
		}

		for _, c := range p.Configurations {
			configurationEntry := model.Configuration{
				ID:             0,
				Characteristic: model.Characteristic{ID: c.CharacteristicID},
			}

			configurationID, err := s.configurations.Store(positionID, &configurationEntry)
			if err != nil {
				return err
			}

			for _, v := range c.VariantIDs {
				variantEntry := model.Variant{ID: v}
				_, err := s.configurations.StoreVariant(configurationID, &variantEntry)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

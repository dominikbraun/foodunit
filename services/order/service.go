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
	orderEntities, err := s.orders.FindByOffer(offerID)

	if err == sql.ErrNoRows {
		return []Order{}, ErrOfferNotFound
	} else if err != nil {
		return []Order{}, err
	}

	orders := make([]Order, 0)

	for _, o := range orderEntities {
		order, err := s.buildOrder(&o)
		if err != nil {
			return []Order{}, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *Service) Get(offerID, userID uint64) (Order, error) {
	orderEntity, err := s.orders.FindByOfferAndUser(offerID, userID)

	if err == sql.ErrNoRows {
		return Order{}, ErrOrderNotFound
	} else if err != nil {
		return Order{}, err
	}

	order, err := s.buildOrder(&orderEntity)

	return order, err
}

func (s *Service) buildOrder(orderEntry *model.Order) (Order, error) {
	order := Order{
		ID:     orderEntry.ID,
		User:   User{Name: orderEntry.User.Name},
		IsPaid: orderEntry.IsPaid,
		Total:  0,
	}

	positionEntities, err := s.positions.FindByOrder(orderEntry.ID)

	if err == sql.ErrNoRows && err != nil {
		return Order{}, err
	}

	for _, p := range positionEntities {
		dishEntity, err := s.dishes.Find(p.Dish.ID)
		alternativeEntity, err := s.dishes.Find(p.Alternative.ID)
		configurations, err := s.loadConfigurations(p.ID)

		if err == sql.ErrNoRows {
			continue
		} else if err != nil {
			return Order{}, err
		}

		position := Position{
			ID: p.ID,
			Dish: Dish{
				ID:    dishEntity.ID,
				Name:  dishEntity.Name,
				Price: dishEntity.Price,
			},
			Alternative: Dish{
				ID:    alternativeEntity.ID,
				Name:  alternativeEntity.Name,
				Price: alternativeEntity.Price,
			},
			Note:           p.Note,
			Configurations: configurations,
		}

		order.Positions = append(order.Positions, position)
		order.Total += dishEntity.Price
	}

	return order, nil
}

func (s *Service) loadConfigurations(positionID uint64) ([]Configuration, error) {
	configurationEntities, err := s.configurations.FindByPosition(positionID)

	if err == sql.ErrNoRows {
		return []Configuration{}, nil
	} else if err != nil {
		return []Configuration{}, err
	}

	configurations := make([]Configuration, 0)

	for _, c := range configurationEntities {
		variantEntities, err := s.configurations.FindVariants(c.ID)

		if err == sql.ErrNoRows {
			continue
		} else if err != nil {
			return []Configuration{}, err
		}

		variants := make([]Variant, 0)

		for _, v := range variantEntities {
			variant := Variant{
				ID:    v.ID,
				Value: v.Value,
				Price: v.Price,
			}

			variants = append(variants, variant)
		}

		configuration := Configuration{
			CharacteristicName: c.Characteristic.Name,
			Multiple:           c.Characteristic.Multiple,
			Variants:           variants,
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

	orderEntity := model.Order{
		User:   model.User{ID: order.UserID},
		IsPaid: false,
	}
	orderID, err := s.orders.Store(order.OfferID, &orderEntity)
	if err != nil {
		return err
	}

	for _, p := range order.Positions {
		positionEntity := model.Position{
			Dish:        model.Dish{ID: p.DishID},
			Alternative: model.Dish{ID: p.AlternativeDishID},
			Note:        p.Note,
		}
		positionID, err := s.positions.Store(orderID, &positionEntity)
		if err != nil {
			return err
		}

		for _, c := range p.Configurations {
			configurationEntity := model.Configuration{
				ID:             0,
				Characteristic: model.Characteristic{ID: c.CharacteristicID},
			}

			configurationID, err := s.configurations.Store(positionID, &configurationEntity)
			if err != nil {
				return err
			}

			for _, v := range c.VariantIDs {
				variantEntity := model.Variant{ID: v}
				_, err := s.configurations.StoreVariant(configurationID, &variantEntity)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

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
	ErrOfferNotFound    = errors.New("the offer could not be found")
	ErrOrderNotFound    = errors.New("the order could not be found")
	ErrActionNotAllowed = errors.New("the action is not allowed")
)

// Service executes order-related business logic and use cases. It is also responsible
// for accessing the model storage under consideration of all business rules.
type Service struct {
	offers          storage.Offer
	orders          storage.Order
	positions       storage.Position
	configurations  storage.Configuration
	dishes          storage.Dish
	characteristics storage.Characteristic
	variants        storage.Variant
}

// NewService creates a new Service instance utilizing the given storage objects.
// The storage objects need to be ready to use for the service.
func NewService(o storage.Offer, odr storage.Order, p storage.Position, c storage.Configuration, d storage.Dish, chr storage.Characteristic, v storage.Variant) *Service {
	service := Service{
		offers:          o,
		orders:          odr,
		positions:       p,
		configurations:  c,
		dishes:          d,
		characteristics: chr,
		variants:        v,
	}
	return &service
}

// GetAll returns all orders for the offer identified by offerID. This will
// also include the individual order positions as well as the user configurations.
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

// Get returns the particular order a user (identified by userID) has made at
// an given offer (identified by offerID).
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

// buildOrder creates an order.Order object out of a model.Order object. This
// includes all order positions and all configurations for each position.
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

// loadConfigurations returns the particular configurations a user has
// created for a given position.
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

// Update updates an user's order. The current updating mechanism does not
// PUT/DELETE/UPDATE a particular order position - instead, the whole order
// will be stored.
// As a consequence, to receive the order, the last stored order is returned.
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

// MarkAsPaid marks an order identified by orderID at an offer identified by
// offerID as paid. The only person allowed to mark an order as paid is the
// offer's owner.
func (s *Service) MarkAsPaid(offerID, orderID, userID uint64) error {
	err := s.orders.MarkAsPaid(orderID)

	ownerID, err := s.offers.OwnerID(offerID)

	if err == sql.ErrNoRows {
		return ErrOfferNotFound
	} else if err != nil {
		return err
	}

	if ownerID != userID {
		return ErrActionNotAllowed
	}

	err = s.orders.MarkAsPaid(orderID)

	if err == sql.ErrNoRows {
		return ErrOrderNotFound
	} else if err != nil {
		return err
	}

	return nil
}

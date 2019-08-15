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

// Package core provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import (
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/pkg/errors"
)

// PositionMap holds multiple positions in an user's order. The key
// represents the DishID, while the value depicts the Note.
type PositionMap map[uint64]string

// OfferInteractor represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type OfferInteractor struct {
	offers    dl.OfferRepository
	orders    dl.OrderRepository
	positions dl.PositionRepository
}

// Creates a new OfferInteractor instance.
func NewOfferInteractor() (*OfferInteractor, error) {
	loader := load.RepositoryLoader

	if !loader.IsReady {
		return nil, errors.New("repository loader has to be initialized first")
	}

	return &OfferInteractor{
		offers:    loader.Offers,
		orders:    loader.Orders,
		positions: loader.Positions,
	}, nil
}

// SaveOrder persists a new order that contains the given positions.
func (o *OfferInteractor) SaveOrder(userID, offerID uint64, pos PositionMap) error {
	order := dl.Order{
		UserID:  userID,
		OfferID: offerID,
	}

	orderID, err := o.orders.Create(&order)
	if err != nil {
		return err
	}

	for id, note := range pos {
		position := dl.Position{
			ID:      0,
			OrderID: orderID,
			DishID:  id,
			Note:    note,
		}
		// ToDo: The error value of Create has to be handled.
		_, _ = o.positions.Create(&position)
	}
	return nil
}

// GetOrders returns all saved orders for a given offer.
func (o *OfferInteractor) GetOrders(offerID uint64) ([]*dl.Order, error) {
	orders, err := o.orders.FindByOfferID(offerID)
	return orders, err
}

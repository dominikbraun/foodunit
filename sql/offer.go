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

// Package sql provides a dl.DataAccess implementation for MariaDB.
package sql

import (
	"github.com/dominikbraun/foodunit/dl"
)

type (
	// offerRepository implements dl.OfferRepository.
	offerRepository struct{}
	// orderRepository implements dl.OrderRepository.
	orderRepository struct{}
	// positionRepository implements dl.PositionRepository.
	positionRepository struct{}
)

// Migrate implements dl.OfferRepository.Migrate.
func (o offerRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.OfferRepository.Create.
func (o offerRepository) Create(offer *dl.Offer) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.OfferRepository.FindByID.
func (o offerRepository) FindByID(id uint64) (dl.Offer, error) {
	panic("implement me")
}

// FindAllActive implements dl.OfferRepository.FindAllActive.
func (o offerRepository) FindAllActive() ([]dl.Offer, error) {
	panic("implement me")
}

// Update implements dl.OfferRepository.Update.
func (o offerRepository) Update(offer *dl.Offer) error {
	panic("implement me")
}

// Delete implements dl.OfferRepository.Delete.
func (o offerRepository) Delete(offer *dl.Offer) error {
	panic("implement me")
}

// Migrate implements dl.OrderRepository.Migrate.
func (o orderRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.OrderRepository.Create.
func (o orderRepository) Create(order *dl.Order) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.OrderRepository.FindByID.
func (o orderRepository) FindByID(id uint64) (dl.Order, error) {
	panic("implement me")
}

// FindByOfferID implements dl.OrderRepository.FindByOfferID.
func (o orderRepository) FindByOfferID(offerID uint64) ([]dl.Order, error) {
	panic("implement me")
}

// Update implements dl.OrderRepository.Update.
func (o orderRepository) Update(order *dl.Order) error {
	panic("implement me")
}

// Delete implements dl.OrderRepository.Delete.
func (o orderRepository) Delete(order *dl.Order) error {
	panic("implement me")
}

// Migrate implements dl.PositionRepository.Migrate.
func (p positionRepository) Migrate() error {
	panic("implement me")
}

// Create implements dl.PositionRepository.Create.
func (p positionRepository) Create(position *dl.Position) (uint64, error) {
	panic("implement me")
}

// FindByID implements dl.PositionRepository.FindByID.
func (p positionRepository) FindByID(id uint64) (dl.Position, error) {
	panic("implement me")
}

// FindByOrderID implements dl.PositionRepository.FindByOrderID.
func (p positionRepository) FindByOrderID(orderID uint64) ([]dl.Position, error) {
	panic("implement me")
}

// Update implements dl.PositionRepository.Update.
func (p positionRepository) Update(position *dl.Position) error {
	panic("implement me")
}

// Delete implements dl.PositionRepository.Delete.
func (p positionRepository) Delete(position *dl.Position) error {
	panic("implement me")
}

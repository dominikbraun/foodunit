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

// Package storage provides storage interfaces and implementations.
package storage

import (
	"github.com/dominikbraun/foodunit/model"
	"time"
)

type Offer interface {
	Entity
	Store(offer *model.Offer) error
	Find(id uint64) (model.Offer, error)
	FindValidFrom(from time.Time) ([]model.Offer, error)
}

type Order interface {
	Entity
	FindByOffer(offerID uint64) ([]model.Order, error)
}

type Position interface {
	Entity
	FindByOrder(orderID uint64) ([]model.Position, error)
}

type Configuration interface {
	Entity
	FindByPosition(positionID uint64) ([]model.Configuration, error)
}

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

// Package storage provides model and session storage implementations.
package storage

import (
	"github.com/dominikbraun/foodunit/dl"
	"time"
)

type Model interface {
	Migrate() error
	Drop() error
}

// RestaurantModel prescribes methods for accessing Restaurant-related data.
type RestaurantModel interface {
	Model
	GetInfo(id uint64) (dl.Restaurant, error)
	Exists(id uint64) (bool, error)
}

type UserModel interface {
	Model
	Create(user dl.User) error
	GetPasswordHash(mailAddr string) ([]byte, error)
	FindByMailAddr(mailAddr string) (dl.User, error)
	MailExists(mailAddr string) (bool, error)
	Exists(id uint64) (bool, error)
}

type CategoryModel interface {
	Model
	FindByRestaurant(restaurantID uint64) ([]dl.Category, error)
}

type DishModel interface {
	Model
	FindByCategory(categoryID uint64) ([]dl.Dish, error)
}

type CharacteristicModel interface {
	Model
}

type VariantModel interface {
	Model
}

type OfferModel interface {
	Model
	Create(offer dl.Offer) error
	GetActive(now time.Time) ([]uint64, error)
}

type OrderModel interface {
	Model
}

type PositionModel interface {
	Model
}

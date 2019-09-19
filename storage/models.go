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

// Storage concludes all Model interfaces. A function or component which would
// normally take the individual interfaces separately as arguments or fields can
// just take a single Storage implementation instead.
type Storage interface {
	RestaurantModel() RestaurantModel
	CategoryModel() CategoryModel
	DishModel() DishModel
	CharacteristicModel() CharacteristicModel
	VariantModel() VariantModel
	UserModel() UserModel
	OfferModel() OfferModel
	OrderModel() OrderModel
	PositionModel() PositionModel
}

// Model prescribes methods for working with meta information such as database structures.
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

// CategoryModel prescribes methods for accessing Category-related data.
type CategoryModel interface {
	Model
	FindByRestaurant(restaurantID uint64) ([]dl.Category, error)
}

// DishModel prescribes methods for accessing Dish-related data.
type DishModel interface {
	Model
	FindByCategory(categoryID uint64) ([]dl.Dish, error)
}

// RestaurantModel prescribes methods for accessing Restaurant-related data.
type CharacteristicModel interface {
	Model
}

// VariantModel prescribes methods for accessing Variant-related data.
type VariantModel interface {
	Model
}

// UserModel prescribes methods for accessing User-related data.
type UserModel interface {
	Model
	Create(user dl.User) error
	GetPasswordHash(mailAddr string) ([]byte, error)
	FindByMailAddr(mailAddr string) (dl.User, error)
	MailExists(mailAddr string) (bool, error)
	Exists(id uint64) (bool, error)
}

// OfferModel prescribes methods for accessing Offer-related data.
type OfferModel interface {
	Model
	Create(offer dl.Offer) error
	GetActive(now time.Time) ([]uint64, error)
}

// OrderModel prescribes methods for accessing Order-related data.
type OrderModel interface {
	Model
}

// PositionModel prescribes methods for accessing Position-related data.
type PositionModel interface {
	Model
}

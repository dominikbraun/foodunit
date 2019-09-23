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

import "github.com/dominikbraun/foodunit/model"

type Restaurant interface {
	Entity
	Find(id uint64) (model.Restaurant, error)
	Exists(id uint64) (bool, error)
}

type Category interface {
	Entity
	FindByRestaurant(restaurantID uint64) ([]model.Category, error)
	Exists(id uint64) (bool, error)
}

type Dish interface {
	Entity
	FindByCategory(categoryID uint64) ([]model.Dish, error)
	Exists(id uint64) (bool, error)
}

type Characteristic interface {
	Entity
	FindByDish(dishID uint64) ([]model.Characteristic, error)
	Exists(id uint64) (bool, error)
}

type Variant interface {
	Entity
	FindByCharacteristic(characteristicID uint64) ([]model.Variant, error)
	Exists(id uint64) (bool, error)
}

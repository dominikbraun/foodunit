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

// Package maria provides storage implementations for MariaDB.
package maria

import (
	"github.com/dominikbraun/foodunit/model"
	"github.com/jmoiron/sqlx"
)

type Dish struct {
	DB *sqlx.DB
}

func NewDish(db *sqlx.DB) *Dish {
	dish := Dish{
		DB: db,
	}
	return &dish
}

func (d *Dish) Prepare() error {
	panic("implement me")
}

func (d *Dish) Drop() error {
	panic("implement me")
}

func (d *Dish) FindByCategory(categoryID uint64) ([]model.Dish, error) {
	panic("implement me")
}

func (d *Dish) Exists(id uint64) (bool, error) {
	panic("implement me")
}

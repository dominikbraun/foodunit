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

func (d *Dish) Create() error {
	query := `
CREATE TABLE dishes (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	description VARCHAR(200) NOT NULL,
	price INTEGER UNSIGNED NOT NULL,
	is_uncertain BOOLEAN NOT NULL,
	is_healthy BOOLEAN NOT NULL,
	is_vegetarian BOOLEAN NOT NULL,
	category_id BIGINT UNSIGNED NOT NULL
)`

	_, err := d.DB.Exec(query)
	return err
}

func (d *Dish) Drop() error {
	query := `DROP TABLE IF EXISTS dishes`
	_, err := d.DB.Exec(query)

	return err
}

func (d *Dish) FindByCategory(categoryID uint64) ([]model.Dish, error) {
	query := `SELECT id, name, description, price, is_uncertain, is_healthy, is_vegetarian FROM dishes WHERE category_id = ?`

	rows, err := d.DB.Queryx(query, categoryID)
	if err != nil {
		return nil, err
	}

	var dishes []model.Dish

	for rows.Next() {
		var dish model.Dish

		err = rows.StructScan(&dish)
		if err != nil {
			return nil, err
		}

		dishes = append(dishes, dish)
	}

	return dishes, nil
}

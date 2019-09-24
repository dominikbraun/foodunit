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

type Category struct {
	DB *sqlx.DB
}

func NewCategory(db *sqlx.DB) *Category {
	category := Category{
		DB: db,
	}
	return &category
}

func (c *Category) Create() error {
	query := `
CREATE TABLE categories (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	restaurant_id BIGINT UNSIGNED NOT NULL
)`

	_, err := c.DB.Exec(query)
	return err
}

func (c *Category) Drop() error {
	query := `DROP TABLE IF EXISTS categories`
	_, err := c.DB.Exec(query)

	return err
}

func (c *Category) FindByRestaurant(restaurantID uint64) ([]model.Category, error) {
	query := `SELECT id, name FROM categories WHERE restaurant_id = ?`

	rows, err := c.DB.Queryx(query, restaurantID)
	if err != nil {
		return nil, err
	}

	var categories []model.Category

	for rows.Next() {
		var category model.Category

		err = rows.StructScan(&category)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

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

// Package mariadb provides MariaDB-compatible model implementations.
package mariadb

import (
	"github.com/dominikbraun/foodunit/dl"
	"github.com/jmoiron/sqlx"
)

// CategoryModel is a storage.CategoryModel implementation.
type CategoryModel struct {
	DB *sqlx.DB
}

// Migrate implements storage.Model.Migrate.
func (c CategoryModel) Migrate() error {
	query := `
CREATE TABLE categories (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	restaurant_id BIGINT UNSIGNED NOT NULL
)`
	_, err := exec(c.DB, query)
	return err
}

// Drop implements storage.Model.Drop.
func (c CategoryModel) Drop() error {
	return drop(c.DB, "categories")
}

// FindByRestaurant implements storage.CategoryModel.FindByRestaurant.
func (c CategoryModel) FindByRestaurant(restaurantID uint64) ([]dl.Category, error) {
	query := `SELECT * FROM categories WHERE restaurant_id = ?`

	rows, err := c.DB.Queryx(query, restaurantID)
	if err != nil {
		return nil, err
	}

	var categories []dl.Category

	for rows.Next() {
		var category dl.Category

		err = rows.StructScan(&category)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

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

// RestaurantModel is a storage.RestaurantModel implementation.
type RestaurantModel struct {
	DB *sqlx.DB
}

// Migrate implements storage.RestaurantModel.Migrate.
func (r RestaurantModel) Migrate() error {
	query := `drop table if exists restaurants`
	_ = r.DB.MustExec(query)

	query = `
CREATE TABLE restaurants (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    postal_code VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    open_mon VARCHAR(50) NOT NULL,
    open_wed VARCHAR(50) NOT NULL,
    open_thu VARCHAR(50) NOT NULL,
    open_fri VARCHAR(50) NOT NULL,
    open_sat VARCHAR(50) NOT NULL,
    open_sun VARCHAR(50) NOT NULL,
    website VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL
)`

	_ = r.DB.MustExec(query)
	return nil
}

// GetInfo implements storage.RestaurantModel.GetInfo.
func (r RestaurantModel) GetInfo(id uint64) (dl.Restaurant, error) {
	query := `SELECT * FROM restaurants WHERE id = ?`

	var restaurant dl.Restaurant
	err := r.DB.QueryRowx(query, id).StructScan(&restaurant)

	return restaurant, err
}

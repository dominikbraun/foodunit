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
	. "github.com/dominikbraun/foodunit/storage/mariadb/query"
	"github.com/jmoiron/sqlx"
)

type RestaurantModel struct {
	db *sqlx.DB
}

func (r RestaurantModel) Migrate() error {
	q := Create("restaurants", Fields{
		"id":          "BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY",
		"name":        "VARCHAR(50)",
		"postal_code": "VARCHAR(50)",
		"city":        "VARCHAR(50)",
		"phone":       "VARCHAR(50)",
		"open_mon":    "VARCHAR(50)",
		"open_tue":    "VARCHAR(50)",
		"open_wed":    "VARCHAR(50)",
		"open_thu":    "VARCHAR(50)",
		"open_fri":    "VARCHAR(50)",
		"open_sat":    "VARCHAR(50)",
		"open_sun":    "VARCHAR(50)",
		"website":     "VARCHAR(50)",
		"is_active":   "BOOLEAN",
	})

	_ = r.db.MustExec(q)
	return nil
}

func (r RestaurantModel) GetInfo(id uint32) (dl.Restaurant, error) {
	q := Select("restaurants", List{"name"}, Conditions{
		"id": "?",
	})

}
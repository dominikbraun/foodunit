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
	"github.com/dominikbraun/foodunit/storage"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// provideDB provides a ready-to-go database connection pool.
func ProvideDBConnection(driver, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "connection failed")
	}

	return db, err
}

// provideRestaurantModel provides a storage.RestaurantModel implementation.
func ProvideRestaurantModel(db *sqlx.DB) storage.RestaurantModel {
	restaurantModel := RestaurantModel{
		DB: db,
	}
	return &restaurantModel
}

// provideRestaurantModel provides a storage.RestaurantModel implementation.
func ProvideUserModel(db *sqlx.DB) storage.UserModel {
	userModel := UserModel{
		DB: db,
	}
	return &userModel
}

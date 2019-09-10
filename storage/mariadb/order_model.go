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
	"github.com/jmoiron/sqlx"
)

// OrderModel is a storage.OrderModel implementation.
type OrderModel struct {
	DB *sqlx.DB
}

// Migrate implements storage.Model.Migrate.
func (c OrderModel) Migrate() error {
	query := `
CREATE TABLE order (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
	is_paid BOOLEAN NOT NULL,
    order_id BIGINT UNSIGNED NOT NULL
)`
	_, err := exec(c.DB, query)
	return err
}

// Drop implements storage.Model.Drop.
func (c OrderModel) Drop() error {
	return drop(c.DB, "order")
}

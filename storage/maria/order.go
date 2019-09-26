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

type Order struct {
	DB *sqlx.DB
}

func NewOrder(db *sqlx.DB) *Order {
	order := Order{
		DB: db,
	}
	return &order
}

func (o *Order) Create() error {
	query := `
CREATE TABLE orders (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	user_id BIGINT UNSIGNED NOT NULL,
	is_paid BOOLEAN NOT NULL,
	offer_id BIGINT UNSIGNED NOT NULL
)`

	_, err := o.DB.Exec(query)
	return err
}

func (o *Order) Drop() error {
	query := `DROP TABLE IF EXISTS orders`
	_, err := o.DB.Exec(query)

	return err
}

func (o *Order) FindByOffer(offerID uint64) ([]model.Order, error) {
	query := `
SELECT o.id, o.is_paid u.id as "user_id.id", u.name as "user_id.name"
FROM offers o
INNER JOIN users u
ON u.id = o.user_id
WHERE o.offer_id = ?`

	rows, err := o.DB.Queryx(query, offerID)
	if err != nil {
		return nil, err
	}

	orders := make([]model.Order, 0)

	for rows.Next() {
		var order model.Order

		if err := rows.StructScan(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

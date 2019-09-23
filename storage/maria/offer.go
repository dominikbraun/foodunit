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

type Offer struct {
	DB *sqlx.DB
}

func NewOffer(db *sqlx.DB) *Offer {
	offer := Offer{
		DB: db,
	}
	return &offer
}

func (o *Offer) Create() error {
	query := `
CREATE TABLE offers (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	owner_user_id BIGINT UNSIGNED NOT NULL,
	restaurant_id BIGINT UNSIGNED NOT NULL,
	valid_from DATETIME NOT NULL,
	valid_to DATETIME NOT NULL,
	responsible_user_id BIGINT UNSIGNED NOT NULL,
	is_placed BOOLEAN NOT NULL,
	ready_at DATETIME NOT NULL,
	paypal_enabled BOOLEAN NOT NULL
)`

	_, err := o.DB.Exec(query)
	return err
}

func (o *Offer) Drop() error {
	query := `DROP TABLE IF EXISTS offers`
	_, err := o.DB.Exec(query)

	return err
}

func (o *Offer) Find(id uint64) (model.Offer, error) {
	panic("implement me")
}

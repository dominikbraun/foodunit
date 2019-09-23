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

type Variant struct {
	DB *sqlx.DB
}

func NewVariant(db *sqlx.DB) *Variant {
	variant := Variant{
		DB: db,
	}
	return &variant
}

func (v *Variant) Create() error {
	query := `
CREATE TABLE variants (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	value VARCHAR(50) NOT NULL,
	is_default BOOLEAN NOT NULL,
	price INTEGER UNSIGNED NOT NULL,
	characteristic_id BIGINT UNSIGNED NOT NULL
)`

	_, err := v.DB.Exec(query)
	return err
}

func (v *Variant) Drop() error {
	query := `DROP TABLE IF EXISTS variants`
	_, err := v.DB.Exec(query)

	return err
}

func (v *Variant) FindByCharacteristic(characteristicID uint64) ([]model.Variant, error) {
	panic("implement me")
}

func (v *Variant) Exists(id uint64) (bool, error) {
	panic("implement me")
}

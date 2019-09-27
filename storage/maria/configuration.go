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

type Configuration struct {
	DB *sqlx.DB
}

func NewConfiguration(db *sqlx.DB) *Configuration {
	configuration := Configuration{
		DB: db,
	}
	return &configuration
}

func (c *Configuration) Create() error {
	query := `
CREATE TABLE configurations (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	position_id BIGINT UNSIGNED NOT NULL,
	characteristic_id BIGINT UNSIGNED NOT NULL
)`
	_, err := c.DB.Exec(query)

	query = `
CREATE TABLE configurations_variants (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	configuration_id BIGINT UNSIGNED NOT NULL,
	variant_id BIGINT UNSIGNED NOT NULL
)`

	return err
}

func (c *Configuration) Drop() error {
	query := `DROP TABLE IF EXISTS configurations, configurations_variants`
	_, err := c.DB.Exec(query)

	return err
}

func (c *Configuration) FindByPosition(positionID uint64) ([]model.Configuration, error) {
	query := `
SELECT conf.id, c.id as "characteristic_id.id", c.name as "characteristic_id.name"
FROM configurations conf
INNER JOIN characteristics c
ON c.id = conf.characteristic_id
WHERE conf.position_id = ?`

	rows, err := c.DB.Queryx(query, positionID)
	if err != nil {
		return nil, err
	}

	configurations := make([]model.Configuration, 0)

	for rows.Next() {
		var configuration model.Configuration

		if err := rows.StructScan(&configuration); err != nil {
			return nil, err
		}

		configurations = append(configurations, configuration)
	}

	return configurations, nil
}

func (c *Configuration) FindVariants(id uint64) ([]model.Variant, error) {
	query := `
SELECT v.id, v.value, v.is_default, v.price
FROM configurations_variants cv
INNER JOIN variants v
ON v.id = cv.variant_id
WHERE cv.configuration_id = ?`

	rows, err := c.DB.Queryx(query, id)
	if err != nil {
		return nil, err
	}

	variants := make([]model.Variant, 0)

	for rows.Next() {
		var variant model.Variant

		if err := rows.StructScan(&variant); err != nil {
			return nil, err
		}
		variants = append(variants, variant)
	}

	return variants, nil
}

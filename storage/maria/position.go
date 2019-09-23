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

type Position struct {
	DB *sqlx.DB
}

func NewPosition(db *sqlx.DB) *Position {
	position := Position{
		DB: db,
	}
	return &position
}

func (p *Position) Create() error {
	query := `
CREATE TABLE positions (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	dish_id BIGINT UNSIGNED NOT NULL,
	alternative_dish_id BIGINT UNSIGNED NOT NULL,
	note VARCHAR(200) NOT NULL,
	order_id BIGINT UNSIGNED NOT NULL
)`
	_, err := p.DB.Exec(query)
	return err
}

func (p *Position) Drop() error {
	query := `DROP TABLE IF EXISTS positions`
	_, err := p.DB.Exec(query)

	return err
}

func (p *Position) FindByOrder(orderID uint64) ([]model.Position, error) {
	panic("implement me")
}

func (p *Position) Exists(id uint64) (bool, error) {
	panic("implement me")
}

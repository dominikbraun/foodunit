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

// Package migration provides tools for performing data migrations.
package migration

import (
	"github.com/dominikbraun/foodunit/storage"
	"github.com/dominikbraun/foodunit/storage/maria"
	"github.com/jmoiron/sqlx"
	"log"
)

type Config struct {
	Driver     string `json:"driver"`
	DSN        string `json:"dsn"`
	DropSchema bool   `json:"drop_schema"`
}

type Manger struct {
	db         *sqlx.DB
	entities   []storage.Entity
	dropSchema bool
}

func NewManager(config *Config) (*Manger, error) {
	m := Manger{}
	var err error

	m.db, err = sqlx.Open(config.Driver, config.DSN)
	m.entities = []storage.Entity{
		maria.NewRestaurant(m.db),
		maria.NewCategory(m.db),
		maria.NewDish(m.db),
		maria.NewCharacteristic(m.db),
		maria.NewVariant(m.db),
		maria.NewUser(m.db),
		maria.NewOffer(m.db),
		maria.NewOrder(m.db),
		maria.NewPosition(m.db),
	}
	m.dropSchema = config.DropSchema

	return &m, err
}

func (m *Manger) RunPreparation() {
	var err error

	if m.dropSchema {
		for _, s := range m.entities {
			if err = s.Drop(); err != nil {
				log.Fatal(err)
			}
		}
	}

	for _, s := range m.entities {
		if err = s.Create(); err != nil {
			log.Fatal(err)
		}
	}
}

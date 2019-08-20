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

// Package sql provides a dl.DataAccess implementation for MariaDB.
package sql

import (
	"fmt"
	"github.com/dominikbraun/foodunit/dl"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const driver string = "mysql"

// Config represents database credentials.
type Config struct {
	Addr   string `json:"host"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbname"`
}

// MariaDB wraps all individual repository implementations and provides
// a method for each repository, meaning that it implements DataAccess.
type MariaDB struct {
	db *sqlx.DB
}

// Open implements dl.DataAccess.Open.
func (m *MariaDB) Open(config interface{}) error {
	c, ok := config.(*Config)
	if !ok {
		return errors.New("config has to be of type *sql.Config")
	}

	dsn := fmt.Sprintf("%s:%s@(%s)/%s", c.User, c.Pass, c.Addr, c.DBName)
	var err error

	if m.db, err = sqlx.Connect(driver, dsn); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// MigrateAll implements dl.DataAccess.MigrateAll.
func (m *MariaDB) MigrateAll() error {
	return nil
}

// Suppliers implements dl.DataAccess.Suppliers.
func (m *MariaDB) Suppliers() dl.SupplierRepository {
	s := supplierRepository{}
	return s
}

// Categories implements dl.DataAccess.Categories.
func (m *MariaDB) Categories() dl.CategoryRepository {
	c := categoryRepository{}
	return c
}

// Dishes implements dl.DataAccess.Dishes.
func (m *MariaDB) Dishes() dl.DishRepository {
	d := dishRepository{}
	return d
}

// Characteristics implements dl.DataAccess.Characteristics.
func (m *MariaDB) Characteristics() dl.CharacteristicRepository {
	c := characteristicRepository{}
	return c
}

// Variants implements dl.DataAccess.Variants.
func (m *MariaDB) Variants() dl.VariantRepository {
	v := variantRepository{}
	return v
}

// Users implements dl.DataAccess.Users.
func (m *MariaDB) Users() dl.UserRepository {
	u := userRepository{}
	return u
}

// Offers implements dl.DataAccess.Offers.
func (m *MariaDB) Offers() dl.OfferRepository {
	o := offerRepository{}
	return o
}

// Orders implements dl.DataAccess.Orders.
func (m *MariaDB) Orders() dl.OrderRepository {
	o := orderRepository{}
	return o
}

// Positions implements dl.DataAccess.Positions.
func (m *MariaDB) Positions() dl.PositionRepository {
	p := positionRepository{}
	return p
}

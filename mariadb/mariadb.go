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

// Package mariadb provides a dl.DataAccess implementation for MariaDB.
package mariadb

import (
	"fmt"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const driver string = "mysql"

type Config struct {
	User   string
	Pass   string
	DBName string
}

type MariaDB struct {
	db *sqlx.DB
}

func (m *MariaDB) Open(config interface{}) error {
	c, ok := config.(*Config)
	if !ok {
		return errors.New("config has to be of type *mariadb.Config")
	}

	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable", c.User, c.DBName)
	var err error

	if m.db, err = sqlx.Connect(driver, dsn); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (m *MariaDB) MigrateAll() error {
	return nil
}

func (m *MariaDB) Offers() dl.OfferRepository {
	panic("implement me")
}

func (m *MariaDB) Orders() dl.OrderRepository {
	panic("implement me")
}

func (m *MariaDB) Positions() dl.PositionRepository {
	panic("implement me")
}

func (m *MariaDB) Suppliers() dl.SupplierRepository {
	panic("implement me")
}

func (m *MariaDB) Categories() dl.CategoryRepository {
	panic("implement me")
}

func (m *MariaDB) Dishes() dl.DishRepository {
	panic("implement me")
}

func (m *MariaDB) Users() dl.UserRepository {
	panic("implement me")
}

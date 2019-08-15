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

// Package mariadb provides repository implementations as SQL gateways.
package mariadb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dominikbraun/foodunit/gateways"
	"github.com/jmoiron/sqlx"
)

// instance represents the database connection to be used by all repositories
// and has to be initialized by the utilizing service with Connect first. This
// object should be exclusively accessed via GetDB().
var instance *sqlx.DB = nil

// Connect establishes a database connection using the given parameters. It
// will initialize the global database instance used by all repositories.
func Connect(c *gateways.Conf) error {
	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable", c.User, c.DBName)

	db, err := sqlx.Connect(c.Driver, dsn)
	if err != nil {
		return err
	}

	instance = db
	return nil
}

// GetDB returns the global database instance. It is recommended to call
// this function instead of accessing the instance directly.
func GetDB() (*sqlx.DB, error) {
	if instance == nil {
		return nil, errors.New("invalid database instance - call `Connect` first")
	}
	return instance, nil
}

// lastInsertID returns the ID of a new inserted row. This function reduces
// repetitive error handling and casts to uint64 in repository implementations.
func lastInsertID(result sql.Result) (uint64, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), err
}

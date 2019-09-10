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

// Package migration provides migration and set up capabilities
package migration

import (
	"github.com/dominikbraun/foodunit/controllers"
	"github.com/dominikbraun/foodunit/storage/mariadb"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Server represents an API server that offers endpoints for data related
// with restaurants, users, offers and orders.
type Server struct {
	controller *controllers.Migration
}

// Setup builds a new Server instance, registers all routes, injects discrete
// model implementations and eventually establishes a database connection.
func Setup(driver, dsn string) (*Server, error) {
	// ToDo: how  to handle  dynamic driver change?
	db, err := mariadb.ProvideDBConnection(driver, dsn)
	if err != nil {
		return nil, err
	}

	migrationController := controllers.ProvideMigrationController(
		mariadb.ProvideRestaurantModel(db),
		mariadb.ProvideUserModel(db),
		mariadb.ProvideCategoryModel(db),
		mariadb.ProvideDishModel(db))

	s := Server{
		controller: migrationController,
	}

	return &s, nil
}

// RunMigration sets up all tables by invoking the individual Migrate() methods.
func (s *Server) Run(drop bool) {
	var err error

	err = s.controller.Migrate(drop)

	if err != nil {
		log.Fatal(err)
	}
}

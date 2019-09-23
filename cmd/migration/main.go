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

// Package main provides a main function for running the migration manager.
package main

import (
	"flag"
	"github.com/dominikbraun/foodunit/migration"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DSNKey            = "dsn"
	DSNDefault        = "root:root@(localhost:3306)/foodunit?parseTime=true"
	DSNHelp           = "The data source name for connecting to the database."
	DropSchemaKey     = "drop-schema"
	DropSchemaDefault = false
	DropSchemaHelp    = "Drops the existing storage schema (e. g. all tables)"
)

func main() {
	config := parseConfig()

	s, err := migration.NewManager(&config)
	if err != nil {
		log.Fatal(err)
	}

	s.RunPreparation()
}

func parseConfig() migration.Config {
	dsn := flag.String(DSNKey, DSNDefault, DSNHelp)
	dropSchema := flag.Bool(DropSchemaKey, DropSchemaDefault, DropSchemaHelp)

	flag.Parse()

	config := migration.Config{
		Driver:     "mysql",
		DSN:        *dsn,
		DropSchema: *dropSchema,
	}

	return config
}

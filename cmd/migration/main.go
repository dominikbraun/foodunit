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

// Package migration provides a runnable server instance which connects to
// a data store and executes all known model migrations.
package main

import (
	"flag"
	"github.com/dominikbraun/foodunit/migration"
	"log"
)

// main sets up a server and invokes its RunMigration method.
func main() {
	dbURI := flag.String("db-uri", "root:root@(localhost:3306)/foodunit", "URI used to connect to the db.")
	drop := flag.Bool("drop", false, "This flag drops all tables before creation")
	flag.Parse()
	s, err := migration.Setup("mysql", *dbURI)
	if err != nil {
		log.Fatal(err)
	}

	s.Run(*drop)
}

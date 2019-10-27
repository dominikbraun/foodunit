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

// Package main provides a main function for running the API server.
package main

import (
	"flag"
	"github.com/dominikbraun/foodunit/server"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DSNKey      = "dsn"
	DSNDefault  = "root:root@(localhost:3306)/foodunit?parseTime=true"
	DSNHelp     = "The data source name for connecting to the database."
	AddrKey     = "addr"
	AddrDefault = ":9292"
	AddrHelp    = "A network address or port the server listens to."
)

// main runs the API server using the provided configuration. This function
// will panic in case an error occurs.
func main() {
	config := parseConfig()

	s, err := server.New(&config)
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}

// parseConfig reads the CLI arguments/flags provided by the user. If no arguments
// are given, the default values will be used instead.
func parseConfig() server.Config {
	dsn := flag.String(DSNKey, DSNDefault, DSNHelp)
	addr := flag.String(AddrKey, AddrDefault, AddrHelp)

	flag.Parse()

	config := server.Config{
		Driver: "mysql",
		DSN:    *dsn,
		Addr:   *addr,
	}

	return config
}

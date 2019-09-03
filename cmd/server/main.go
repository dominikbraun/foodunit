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

// Package server provides a runnable server instance which exposes a REST API.
package main

import (
	"log"

	"github.com/dominikbraun/foodunit/server"
)

// main sets up a server and invokes its Run method.
func main() {
	s, err := server.Setup("mysql", "root:root@(localhost:3306)/foodunit")
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}

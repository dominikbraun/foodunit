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

// Package controllers provides various controllers for triggering the core logic.
package controllers

import (
	"github.com/dominikbraun/foodunit/storage"
)

// Migration represents the handler for the initial set up of FootUnit
type Migration struct {
	Restaurants storage.RestaurantModel
	Users       storage.UserModel
}

func (m *Migration) Migrate(drop bool) error {
	if drop {
		// ToDo
	}

	m.Restaurants.Migrate()
	m.Users.Migrate()

	return nil
}

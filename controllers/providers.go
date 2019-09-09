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

import "github.com/dominikbraun/foodunit/storage"

// provideRESTController provides a controller instance for handing REST requests.
func ProvideRESTController(r storage.RestaurantModel, u storage.UserModel) *REST {
	controller := REST{
		Restaurants: r,
		Users:       u,
	}
	return &controller
}

// provideMigrationController provides a controller instance for handling the initial FoodUnit set up.
func ProvideMigrationController(r storage.RestaurantModel, u storage.UserModel) *Migration {
	controller := Migration{
		Restaurants: r,
		Users:       u,
	}
	return &controller
}

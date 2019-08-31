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

// Package handlers provides various handlers for triggering the core logic.
package handlers

import (
	"github.com/dominikbraun/foodunit/core"
	"github.com/dominikbraun/foodunit/storage"
	"net/http"
	"strconv"
)

// Renderer represents a type which writes any struct as a JSON object.
type Renderer interface {
	JSON(http.ResponseWriter, *http.Request, interface{})
}

// REST represents the common handler for REST API requests.
type REST struct {
	render      Renderer
	restaurants storage.RestaurantModel
}

// GetRestaurantInfo is responsible for calling core.GetRestaurantInfo.
func (rest *REST) GetRestaurantInfo(w http.ResponseWriter, r *http.Request) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.Header.Get("id"))
		if err != nil {
			// ToDo: Handle type error properly
			return
		}
		restaurantInfo, err := core.GetRestaurantInfo(uint32(id), rest.restaurants)
		if err != nil {
			// ToDo: Handle model error properly
			return
		}
		rest.render.JSON(w, r, restaurantInfo)
	}
	return http.HandlerFunc(fn)
}

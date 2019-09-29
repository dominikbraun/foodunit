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

// Package rest provides a controller for handling API requests.
package rest

import (
	"github.com/dominikbraun/foodunit/services/restaurant"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (c *Controller) RestaurantInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		info, err := c.restaurantService.Info(uint64(id))

		if err != nil && err == restaurant.ErrRestaurantNotFound {
			respond(w, r, http.StatusNotFound, err.Error())
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, info)
		return
	}
}

func (c *Controller) RestaurantMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		menu, err := c.restaurantService.Menu(uint64(id))

		if err != nil && (err == restaurant.ErrRestaurantNotFound || err == restaurant.ErrMenuNotFound) {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, menu)
		return
	}
}

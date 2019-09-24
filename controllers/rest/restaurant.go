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
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (c *Controller) RestaurantInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidNumberFormat.Error())
			return
		}

		info, err := c.restaurantService.Info(uint64(id))

		if err == restaurant.ErrRestaurantNotFound {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, restaurant.ErrRestaurantNotFound.Error())
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, ErrProcessingFailed.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, info)
		return
	}
}

func (c *Controller) RestaurantMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidNumberFormat.Error())
			return
		}

		menu, err := c.restaurantService.Menu(uint64(id))

		if err == restaurant.ErrRestaurantNotFound || err == restaurant.ErrMenuNotFound {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, err)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, ErrProcessingFailed.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, menu)
		return
	}
}

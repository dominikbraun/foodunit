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

// Package api provides controller functions for handling API requests.
package api

import (
	"github.com/dominikbraun/foodunit/services/restaurant"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

var (
	ErrInvalidNumberFormat = errors.New("provided number format is not valid")
	ErrProcessingFailed    = errors.New("failed processing the request")
)

func RestaurantInfo(service *restaurant.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidNumberFormat.Error())
			return
		}

		info, err := service.Info(uint64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, ErrProcessingFailed.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, info)
		return
	}
}

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
	"github.com/dominikbraun/foodunit/services/dish"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (c *Controller) GetCharacteristics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dishID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		characteristics, err := c.dishService.GetCharacteristics(uint64(dishID))

		if err != nil && err == dish.ErrCharacteristicsNotFound {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, characteristics)
		return
	}
}

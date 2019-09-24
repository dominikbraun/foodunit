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
	"encoding/json"
	"github.com/dominikbraun/foodunit/services/user"
	"github.com/go-chi/render"
	"net/http"
)

func (c *Controller) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var registration user.Registration
		err := json.NewDecoder(r.Body).Decode(&registration)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidFormData.Error())
			return
		}

		success, err := c.userService.Register(&registration)

		if err != nil {
			if err == user.ErrUserExists {
				w.WriteHeader(http.StatusUnprocessableEntity)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			render.JSON(w, r, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, success)
		return
	}
}
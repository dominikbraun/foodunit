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
	"github.com/dominikbraun/foodunit/services/order"
	"github.com/dominikbraun/foodunit/session"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (c *Controller) AllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offerID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidNumberFormat.Error())
			return
		}

		orders, err := c.orderService.GetAll(uint64(offerID))

		if err == order.ErrOfferNotFound {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, order.ErrOfferNotFound.Error())
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, ErrProcessingFailed.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, orders)
		return
	}
}

func (c *Controller) GetOrder(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offerID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidNumberFormat.Error())
			return
		}

		userID, ok := session.Get(r.Context(), "uid").(uint64)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			render.JSON(w, r, ErrForbiddenAction.Error())
			return
		}

		userOrder, err := c.orderService.Get(uint64(offerID), userID)

		if err != nil {
			if err == order.ErrOfferNotFound || err == order.ErrOrderNotFound {
				w.WriteHeader(http.StatusNotFound)
				render.JSON(w, r, err.Error())
				return
			} else {
				// ToDo: Handle all 500's like this (printing ErrProcessingFailed)
				w.WriteHeader(http.StatusInternalServerError)
				render.JSON(w, r, ErrProcessingFailed.Error())
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, userOrder)
		return
	}
}

func (c *Controller) UpdateOrder(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var update order.Update
		err := json.NewDecoder(r.Body).Decode(&update)
		offerID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidNumberFormat.Error())
			return
		}

		userID, ok := session.Get(r.Context(), "uid").(uint64)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			render.JSON(w, r, ErrForbiddenAction.Error())
			return
		}

		update.OfferID = uint64(offerID)
		update.UserID = userID

		err = c.orderService.Update(&update)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, ErrProcessingFailed.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, true)
		return
	}
}

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
	"github.com/dominikbraun/foodunit/services/offer"
	"github.com/dominikbraun/foodunit/session"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (c *Controller) CreateOffer(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creation offer.Creation
		err := json.NewDecoder(r.Body).Decode(&creation)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidFormData.Error())
			return
		}

		uid, ok := session.Get(r.Context(), "uid").(uint64)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			render.JSON(w, r, ErrForbiddenAction.Error())
			return
		}

		err = c.offerService.Create(&creation, uid)

		if err != nil {
			if err == offer.ErrUserNotFound || err == offer.ErrRestaurantNotFound {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			render.JSON(w, r, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, true)
		return
	}
}

func (c *Controller) ActiveOffers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offers, err := c.offerService.Active()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, ErrProcessingFailed.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, offers)
		return
	}
}

func (c *Controller) GetOffer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			render.JSON(w, r, ErrInvalidNumberFormat.Error())
			return
		}

		offerView, err := c.offerService.Get(uint64(id))

		if err == offer.ErrOfferNotFound {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, offer.ErrOfferNotFound.Error())
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, offerView)
		return
	}
}

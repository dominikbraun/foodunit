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
	"net/http"
	"strconv"
)

func (c *Controller) CreateOffer(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creation offer.Creation
		err := json.NewDecoder(r.Body).Decode(&creation)

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidFormData.Error())
			return
		}

		uid, ok := session.Get(r.Context(), "uid").(uint64)
		if !ok {
			respond(w, r, http.StatusForbidden, ErrForbiddenAction.Error())
			return
		}

		err = c.offerService.Create(&creation, uid)

		if err != nil && (err == offer.ErrUserNotFound || err == offer.ErrRestaurantNotFound) {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

func (c *Controller) ActiveOffers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offers, err := c.offerService.Active()

		if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, offers)
		return
	}
}

func (c *Controller) GetOffer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		offerView, err := c.offerService.Get(uint64(id))

		if err != nil && err == offer.ErrOfferNotFound {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, offerView)
		return
	}
}

func (c *Controller) CancelOffer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		err = c.offerService.Cancel(uint64(id))

		if err != nil && err == offer.ErrOfferNotFound {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

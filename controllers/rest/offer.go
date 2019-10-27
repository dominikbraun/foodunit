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
	"github.com/dominikbraun/foodunit/services/order"
	"github.com/dominikbraun/foodunit/session"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// CreateOffer is responsible for creating a new offer. All user-related data
// needs to be provided by the accepted session manager.
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

// ActiveOffers is responsible for retrieving all active offers.
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

// OldOffers is responsible for retrieving all expired offers.
func (c *Controller) OldOffers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offers, err := c.offerService.Old()

		if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, offers)
		return
	}
}

// GetOffers is responsible for retrieving all required data for a given offer.
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

// CancelOffer is responsible for any offer cancellations. All user-related data
// needs to be provided by the accepted session manager.
func (c *Controller) CancelOffer(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offerID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		userID, ok := session.Get(r.Context(), "uid").(uint64)
		if !ok {
			respond(w, r, http.StatusForbidden, ErrForbiddenAction.Error())
			return
		}

		err = c.offerService.Cancel(uint64(offerID), userID)

		if err != nil && (err == offer.ErrOfferNotFound || err == offer.ErrRestaurantNotFound) {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil && err == offer.ErrActionNotAllowed {
			respond(w, r, http.StatusForbidden, ErrForbiddenAction.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

// MarkOrderAsPaid is responsible for setting the IsPaid property of a given order
// to true. All user-related data needs to be provided by the accepted session manager.
func (c *Controller) MarkOrderAsPaid(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offerID, err := strconv.Atoi(chi.URLParam(r, "id"))
		orderID, err := strconv.Atoi(chi.URLParam(r, "orderID"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		userID, ok := session.Get(r.Context(), "uid").(uint64)
		if !ok {
			respond(w, r, http.StatusForbidden, ErrForbiddenAction.Error())
			return
		}

		err = c.orderService.MarkAsPaid(uint64(offerID), uint64(orderID), userID)

		if err != nil && (err == order.ErrOfferNotFound || err == order.ErrOrderNotFound) {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil && err == order.ErrActionNotAllowed {
			respond(w, r, http.StatusForbidden, ErrForbiddenAction.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

// SetReadyAt is responsible for setting the IsReady property of a given offer.
// All user-related data needs to be provided by the accepted session manager.
func (c *Controller) SetReadyAt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var setter offer.ReadyAtSetter
		err := json.NewDecoder(r.Body).Decode(&setter)

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		err = c.offerService.SetReadyAt(uint64(id), setter)

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

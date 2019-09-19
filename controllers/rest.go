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

// Package controllers provides various controllers for triggering the core logic.
package controllers

import (
	"encoding/json"
	"github.com/dominikbraun/foodunit/core/dto"
	"github.com/dominikbraun/foodunit/session"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"

	"github.com/dominikbraun/foodunit/core"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/go-chi/render"
)

// REST represents the common handler for REST API requests.
type REST struct {
	Manager     session.Manager
	Restaurants storage.RestaurantModel
	Categories  storage.CategoryModel
	Dishes      storage.DishModel
	Users       storage.UserModel
	Offers      storage.OfferModel
}

// GetRestaurantInfo is responsible for invoking core.GetRestaurantInfo.
func (rest *REST) GetRestaurantInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// ToDo: Handle type error properly
		return
	}
	restaurantInfo, err := core.GetRestaurantInfo(uint64(id), rest.Restaurants)
	if err != nil {
		// ToDo: Handle core error properly
		return
	}
	render.JSON(w, r, restaurantInfo)
}

// RegisterUser is responsible for invoking core.RegisterUser.
func (rest *REST) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registration dto.UserRegistration
	err := json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	_ = r.Body.Close()

	err = core.RegisterUser(registration, rest.Users)
	if err != nil {
		// ToDo: Handle core error properly
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, true)
}

// Authenticate logs in the authenticating user by comparing the provided password
// with the stored password and creating a new session if the comparison was successful.
func (rest *REST) Login(w http.ResponseWriter, r *http.Request) {
	var login dto.UserLogin
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	_ = r.Body.Close()

	success, err := core.Authenticate(login, rest.Users)
	if err != nil {
		// ToDo: Handle core error properly
		render.JSON(w, r, err.Error())
		return
	}

	if success {
		rest.Manager.Put(r.Context(), "logged_in", true)
	}

	render.JSON(w, r, success)
}

// CreateOffer stores a new user-created offer.
func (rest *REST) CreateOffer(w http.ResponseWriter, r *http.Request) {
	var offerJson struct {
		Restaurant    uint64 `json:"restaurant_id"`
		ValidFrom     string `json:"valid_from"`
		ValidTo       string `json:"valid_to"`
		Responsible   uint64 `json:"responsible_user_id"`
		IsPlaced      bool   `json:"is_placed"`
		ReadyAt       string `json:"ready_at"`
		PaypalEnabled bool   `json:"paypal_enabled"`
	}

	err := json.NewDecoder(r.Body).Decode(&offerJson)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	_ = r.Body.Close()

	dates, err := parseDates(offerJson.ValidFrom, offerJson.ValidTo, offerJson.ReadyAt)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	offer := dto.NewOffer{
		Restaurant:    offerJson.Restaurant,
		ValidFrom:     dates[0],
		ValidTo:       dates[1],
		Responsible:   offerJson.Responsible,
		IsPlaced:      offerJson.IsPlaced,
		ReadyAt:       dates[2],
		PaypalEnabled: offerJson.PaypalEnabled,
	}

	err = core.CreateOffer(offer, rest.Offers, rest.Users, rest.Restaurants)
	if err != nil {
		// ToDo: Handle core error properly
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, true)
}

// GetActiveOffers returns the ids for all offers currently active
func (rest *REST) GetActiveOffers(w http.ResponseWriter, r *http.Request) {
	offers, err := core.GetActiveOffers(rest.Offers)

	if err != nil {
		// ToDo: Handle core error properly
		return
	}
	render.JSON(w, r, offers)
}

// GetRestaurantMenu invokes core.GetRestaurantMenu in order to retrieve a restaurant's menu.
func (rest *REST) GetRestaurantMenu(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// ToDo: Handle type error properly
		return
	}

	menu, err := core.GetRestaurantMenu(uint64(id), rest.Categories, rest.Dishes)
	if err != nil {
		// ToDo: Handle core error properly
		render.JSON(w, r, err.Error())
	}

	render.JSON(w, r, menu)
}

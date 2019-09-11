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
	"github.com/go-chi/chi"
	"net/http"
	"strconv"

	"github.com/dominikbraun/foodunit/core"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/go-chi/render"
)

// REST represents the common handler for REST API requests.
type REST struct {
	Restaurants storage.RestaurantModel
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

func (rest *REST) Authenticate(w http.ResponseWriter, r *http.Request) {
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

	render.JSON(w, r, success)
}

func (rest *REST) CreateOffer(w http.ResponseWriter, r *http.Request) {
	var offer dto.NewOffer
	err := json.NewDecoder(r.Body).Decode(&offer)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	_ = r.Body.Close()

	err = core.CreateOffer(offer, rest.Offers, rest.Users, rest.Restaurants)
	if err != nil {
		// ToDo: Handle core error properly
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, true)
}

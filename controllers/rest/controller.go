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
	"github.com/dominikbraun/foodunit/services/offer"
	"github.com/dominikbraun/foodunit/services/order"
	"github.com/dominikbraun/foodunit/services/restaurant"
	"github.com/dominikbraun/foodunit/services/user"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrInvalidNumberFormat = errors.New("the provided number format is not valid")
	ErrProcessingFailed    = errors.New("failed processing the request")
	ErrInvalidFormData     = errors.New("the provided form data is not valid")
	ErrForbiddenAction     = errors.New("no permission for performing the action")
)

type Controller struct {
	restaurantService *restaurant.Service
	dishService       *dish.Service
	userService       *user.Service
	offerService      *offer.Service
	orderService      *order.Service
}

func NewController(r *restaurant.Service, d *dish.Service, u *user.Service, o *offer.Service, odr *order.Service) *Controller {
	controller := Controller{
		restaurantService: r,
		dishService:       d,
		userService:       u,
		offerService:      o,
		orderService:      odr,
	}
	return &controller
}

func respond(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	w.WriteHeader(status)
	render.JSON(w, r, v)
}

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
	"github.com/dominikbraun/foodunit/session"
	"net/http"
)

func (c *Controller) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var registration user.Registration
		err := json.NewDecoder(r.Body).Decode(&registration)

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidFormData.Error())
			return
		}

		success, err := c.userService.Register(&registration)

		if err != nil && (err == user.ErrUserExists || err == user.ErrPasswordInvalid || err == user.ErrConfirmationMailNotSent) {
			respond(w, r, http.StatusUnprocessableEntity, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, success)
		return
	}
}

func (c *Controller) Login(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login user.Login
		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidFormData.Error())
			return
		}

		uid, err := c.userService.Authenticate(&login)

		if err != nil && (err == user.ErrPasswordIncorrect || err == user.ErrUserNotFound) {
			respond(w, r, http.StatusUnauthorized, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		if uid != 0 {
			session.Put(r.Context(), "authenticated", true)
			session.Put(r.Context(), "uid", uid)
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

func (c *Controller) Logout(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session.Remove(r.Context(), "authenticated")
		session.Remove(r.Context(), "uid")

		respond(w, r, http.StatusOK, true)
		return
	}
}

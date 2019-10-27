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
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// RegisterUser is responsible for triggering an user registration.
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

// Login is responsible for logging in an existing user. The provided session manager
// will create a new user session if the login has been successful.
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
			// send false, not http.StatusUnauthorized as we should not control the logic by errors.
			// wrong user / password is not a program-error.
			respond(w, r, http.StatusOK, false)
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

// IsAuthenticated is responsible for determining whether a user is logged in or not.
// All user-related data needs to be provided by the accepted session manager.
func (c *Controller) IsAuthenticated(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := session.Get(r.Context(), "authenticated").(bool)

		respond(w, r, http.StatusOK, ok)
		return
	}
}

// ConfirmMailAddr is responsible for confirming a user's mail address who has just
// registered. The corresponding confirmation mail has been sent by RegisterUser.
func (c *Controller) ConfirmMailAddr() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")
		if token == "" {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidFormData.Error())
			return
		}

		err := c.userService.ConfirmMailAddr(token)

		if err != nil && err == user.ErrTokenInvalid {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidFormData.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

// Logout is responsible for quitting a user session. The provided session manager will
// remove any data associated with the particular session.
func (c *Controller) Logout(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session.Remove(r.Context(), "authenticated")
		session.Remove(r.Context(), "uid")

		respond(w, r, http.StatusOK, true)
		return
	}
}

// SetPaypalMailAddr is responsible for setting the PaypayMailAddr property of a given user.
func (c *Controller) SetPaypalMailAddr(session session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var setter user.PaypalMailAddrSetter
		err := json.NewDecoder(r.Body).Decode(&setter)

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidFormData.Error())
			return
		}

		id, ok := session.Get(r.Context(), "uid").(uint64)
		if !ok {
			respond(w, r, http.StatusForbidden, ErrForbiddenAction.Error())
			return
		}

		err = c.userService.SetPaypalMailAddr(id, setter)

		if err != nil && err == user.ErrUserNotFound {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
		}

		respond(w, r, http.StatusOK, true)
		return
	}
}

// GetUser is responsible for retrieving all data for a given user.
func (c *Controller) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			respond(w, r, http.StatusUnprocessableEntity, ErrInvalidNumberFormat.Error())
			return
		}

		publicUser, err := c.userService.Get(uint64(id))

		if err != nil && err == user.ErrUserNotFound {
			respond(w, r, http.StatusNotFound, err.Error())
			return
		} else if err != nil {
			respond(w, r, http.StatusInternalServerError, ErrProcessingFailed.Error())
			return
		}

		respond(w, r, http.StatusOK, publicUser)
		return
	}
}

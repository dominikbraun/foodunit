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

// Package middleware provides request handling middleware to be used by a router.
package middleware

import (
	"github.com/dominikbraun/foodunit/session"
	"github.com/go-chi/render"
	"net/http"
)

// Middleware is a function returning an HTTP handler function which takes another
// handler function that will be executed subsequently.
type Middleware func(http.Handler) http.Handler

// Authenticate returns a middleware that provides a handler for checking if a
// `mail_addr` exists in the current session using a SessionManager instance.
func Authenticate(manager session.Manager) Middleware {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mailAddr := manager.Get(r.Context(), "mail_addr")

			if mailAddr == "" {
				render.JSON(w, r, "Forbidden")
			}
			next.ServeHTTP(w, r)
		})
	}
}

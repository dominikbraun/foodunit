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

// Package middleware provides HTTP middleware for request handling chains.
package middleware

import (
	"github.com/dominikbraun/foodunit/session"
	"github.com/go-chi/render"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Authenticate(session session.Manager) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authenticated := session.GetBool(r.Context(), "authenticated")

			if !authenticated {
				w.WriteHeader(http.StatusForbidden)
				render.JSON(w, r, "access forbidden")
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

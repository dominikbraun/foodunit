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

// Package server provides functions to control FoodUnit's HTTP Server.
package server

import (
	"github.com/dominikbraun/foodunit/controllers/http"
	"github.com/go-chi/chi"
)

// mountRoutes builds all API routes and registers the responsible handlers.
func (srv *Server) mountRoutes() {
	r := chi.NewRouter()

	r.Route("/offers", func(r chi.Router) {
		r.Post("/new", http.PostOffer)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.GetOffer)
			r.Put("/", http.PutOffer)
			r.Delete("/", http.DeleteOffer)

			r.Route("/orders", func(r chi.Router) {
				r.Get("/", http.GetOrders)
			})

			r.Route("/order", func(r chi.Router) {
				r.Post("/", http.F)
				r.Get("/", http.F)
				r.Put("/", http.F)
			})
		})
	})

	r.Route("/suppliers", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.F)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", http.F)
		r.Post("/login", http.F)
		r.Get("/logout", http.F)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", http.F)
		})
	})

	srv.router.Mount("/", r)
}

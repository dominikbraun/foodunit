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

// Package server provides an executable server which exposes a REST API.
package server

import (
	"github.com/dominikbraun/foodunit/controllers/api"
	"github.com/go-chi/chi"
)

func (s *Server) mountRoutes() {
	r := chi.NewRouter()

	r.Use(s.session.LoadAndSave)

	r.Route("/restaurants", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/info", api.RestaurantInfo(s.restaurantService))
			r.Get("/menu", nil)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", nil)
		r.Post("/login", nil)
		r.Get("/logout", nil)
	})

	r.Route("/offers", func(r chi.Router) {
		r.Post("/create", nil)
		r.Get("/active", nil)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", nil)

			r.Route("/orders", func(r chi.Router) {
				r.Get("/all", nil)
				r.Get("/mine", nil)
				r.Post("/mine", nil)
			})
		})
	})

	r.Route("/dishes", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/characteristics", nil)
		})
	})

	s.router.Mount("/v1", r)
}

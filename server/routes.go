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

// Package server provides a server which exposes a REST API.
package server

import (
	"github.com/dominikbraun/foodunit/middleware"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// mountRoutes builds all supported routes, registers the corresponding controllers
// and mounts that routes to the server's router instance.
func (s *Server) mountRoutes() {
	r := chi.NewRouter()

	r.Use(s.manager.LoadAndSave)

	r.Route("/restaurants", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			// Get meta information for the restaurant.
			r.Get("/info", s.controller.GetRestaurantInfo)
			// Get the restaurant's menu.
			r.Get("/menu", nil)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", s.controller.RegisterUser)
		r.Post("/login", s.controller.Login)
		r.Get("/logout", nil)

		r.Route("/{id}", func(r chi.Router) {
			// Get the user's profile data.
			r.Get("/profile", nil)
		})
	})

	r.With(middleware.Authenticate(s.manager)).
		Route("/offers", func(r chi.Router) {
			// Create a new offer.
			r.Post("/create", s.controller.CreateOffer)
			// Get all active offers.
			r.Get("/active", nil)

			r.Route("/{id}", func(r chi.Router) {
				// Get information for the offer.
				r.Get("/", nil)

				r.Route("/orders", func(r chi.Router) {
					// Get all orders for the offer.
					r.Get("/all", nil)
					// Get the user's order for the offer.
					r.Get("/mine", nil)
					// Save the user's order.
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

// useReverseProxy sets up a single host reverse proxy prevent cross origin errors.
func (s *Server) useReverseProxy(clientURL string) error {
	client, err := url.Parse(clientURL)
	if err != nil {
		return errors.Wrap(err, "client URL is not valid")
	}

	s.router.NotFound(func(res http.ResponseWriter, req *http.Request) {
		httputil.NewSingleHostReverseProxy(client).ServeHTTP(res, req)
	})

	return nil
}

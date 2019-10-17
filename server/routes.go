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
	"github.com/dominikbraun/foodunit/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (s *Server) mountRoutes() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Use(s.session.LoadAndSave)

	r.With(middleware.Authenticate(s.session)).
		Route("/restaurants", func(r chi.Router) {
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/info", s.controller.RestaurantInfo())
				r.Get("/menu", s.controller.RestaurantMenu())
			})
		})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", s.controller.RegisterUser())
		r.Post("/login", s.controller.Login(s.session))

		r.With(middleware.Authenticate(s.session)).
			Post("/paypal-mail-addr", s.controller.SetPaypalMailAddr(s.session))

		r.Get("/confirm/{token}", s.controller.ConfirmMailAddr())
		r.Get("/logout", s.controller.Logout(s.session))
	})

	r.With(middleware.Authenticate(s.session)).
		Route("/offers", func(r chi.Router) {
			r.Post("/create", s.controller.CreateOffer(s.session))
			r.Get("/active", s.controller.ActiveOffers())

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", s.controller.GetOffer())
				r.Get("/cancel", s.controller.CancelOffer(s.session))
				r.Post("/ready-at", s.controller.SetReadyAt())

				r.Route("/orders", func(r chi.Router) {
					r.Get("/all", s.controller.AllOrders())
					r.Get("/mine", s.controller.GetOrder(s.session))
					r.Post("/mine", s.controller.UpdateOrder(s.session))

					r.Route("/{orderID}", func(r chi.Router) {
						r.Get("/mark-as-paid", s.controller.MarkOrderAsPaid(s.session))
					})
				})
			})
		})

	r.Route("/dishes", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/characteristics", s.controller.GetCharacteristics())
		})
	})

	s.router.Mount("/v1", r)
}

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
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// newRouter creates a route multiplexer and registers essential middleware.
func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)
	return r
}

// mountRoutes builds all supported routes, registers the corresponding handlers
// and mounts that routes to the server's router instance.
func (s *Server) mountRoutes() {
	r := newRouter()

	r.Route("/restaurants", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/info", s.rest.GetRestaurantInfo)
		})
	})

	s.router.Mount("/v1", r)
}

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

	r.Route("/restaurants", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/info", s.controller.GetRestaurantInfo)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", s.controller.RegisterUser)
	})

	s.router.Mount("/v1", r)
}

// useReverseProxy sets up a single host reverse proxy to
func (fs *Server) useReverseProxy(clientURL string) error {
	client, err := url.Parse(clientURL)
	if err != nil {
		return errors.Wrap(err, "client URL is not valid")
	}

	fs.router.NotFound(func(res http.ResponseWriter, req *http.Request) {
		httputil.NewSingleHostReverseProxy(client).ServeHTTP(res, req)
	})

	return nil
}

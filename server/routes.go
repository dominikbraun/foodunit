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
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
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

	s.router.Mount("/v1", r)
}

// Setup fileserver to serve the client files.
func (s *Server) mountClientRoutes(clientURL string) {
	//r := chi.NewRouter()

	if clientURL != "" {
		// Setup proxy which forwards everything to the client. Needed to prevent cross origin error if the client is provided
		// by another Server. (e.g. by yarn serve in development environment)

		client, err := url.Parse(clientURL)
		if err == nil {
			// EVERYTHING redirect to client
			s.router.NotFound(func(res http.ResponseWriter, req *http.Request) {
				httputil.NewSingleHostReverseProxy(client).ServeHTTP(res, req)
			})
		} else {
			log.Fatal("the provided client url is not a valid url: ", clientURL, err)
		}

		return
	}

	folder := "client/dist"

	// Handle static content, we have to explicitly put our top level dirs in here
	// - otherwise the NotFoundHandler will catch them
	fileServer(s.router, "/js", http.Dir(folder+"/js"))
	fileServer(s.router, "/css", http.Dir(folder+"/css"))
	fileServer(s.router, "/img", http.Dir(folder+"/img"))

	// EVERYTHING else redirect to index.html
	s.router.NotFound(func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, folder+"/index.html")
	})
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

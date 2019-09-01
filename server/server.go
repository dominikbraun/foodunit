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
	"context"
	"github.com/dominikbraun/foodunit/handlers"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server represents an API server that offers endpoints for data related
// with restaurants, users, offers and orders.
type Server struct {
	*http.Server
	router    *chi.Mux
	handler   handlers.REST
	interrupt chan os.Signal
}

// New creates a Server instance and returns a reference to it.
func New() *Server {
	s := Server{
		Server:    nil,
		router:    newRouter(),
		handler:   handlers.REST{},
		interrupt: make(chan os.Signal),
	}
	s.Server = &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	return &s
}

// Run mounts all API routes, establishes a database connection and starts
// listening to the specified port. The server can be shut down with Ctrl + C.
func (s *Server) Run() {
	s.mountRoutes()
	signal.Notify(s.interrupt, os.Interrupt)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	<-s.interrupt
	c, cancel := context.WithTimeout(context.Background(), time.Second*5)

	if err := s.Shutdown(c); err != nil {
		log.Println(err)
	}
	defer cancel()
}

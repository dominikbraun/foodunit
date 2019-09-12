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
	"github.com/dominikbraun/foodunit/session"
	"github.com/dominikbraun/foodunit/storage/mariadb"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"

	"github.com/dominikbraun/foodunit/controllers"
	_ "github.com/go-sql-driver/mysql"
)

// Server represents an API server that offers endpoints for data related
// with restaurants, users, offers and orders.
type Server struct {
	*http.Server
	router     *chi.Mux
	controller *controllers.REST
	interrupt  chan os.Signal
	manager    session.Manager
}

// Setup builds a new Server instance, registers all routes, injects discrete
// model implementations and eventually establishes a database connection.
func Setup(driver, dsn string, clientURL string) (*Server, error) {
	// ToDo: how  to handle  dynamic driver change?
	db, err := mariadb.ProvideDBConnection(driver, dsn)
	if err != nil {
		return nil, err
	}

	router := provideRouter()

	restaurantModel := mariadb.ProvideRestaurantModel(db)
	userModel := mariadb.ProvideUserModel(db)
	offerModel := mariadb.ProvideOfferModel(db)
	restController := controllers.ProvideRESTController(restaurantModel, userModel, offerModel)

	storage := mariadb.ProvideSessionStorage(db)
	manager := session.ProvideManager(storage)

	s := Server{
		Server: &http.Server{
			Addr:    ":9292",
			Handler: router,
		},
		router:     router,
		controller: restController,
		interrupt:  make(chan os.Signal),
		manager:    manager,
	}

	s.mountRoutes()

	if clientURL != "" {
		if err := s.useReverseProxy(clientURL); err != nil {
			return nil, err
		}
		return &s, nil
	}

	signal.Notify(s.interrupt, os.Interrupt)

	return &s, nil
}

// Run mounts all API routes, establishes a database connection and starts
// listening to the specified port. The server can be shut down with Ctrl + C.
func (s *Server) Run() {
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	<-s.interrupt
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)

	if err := s.Shutdown(timeout); err != nil {
		log.Println(err)
	}

	defer cancel()
}

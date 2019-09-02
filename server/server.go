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
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dominikbraun/foodunit/handlers"
	"github.com/dominikbraun/foodunit/storage/mariadb"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var db *sqlx.DB

// Server represents an API server that offers endpoints for data related
// with restaurants, users, offers and orders.
type Server struct {
	*http.Server
	router    *chi.Mux
	rest      handlers.REST
	interrupt chan os.Signal
}

// Setup builds a new Server instance, registers all routes, injects discrete
// model implementations and eventually establishes a database connection.
func Setup(driver, dsn string) (*Server, error) {
	s := Server{
		router:    newRouter(),
		interrupt: make(chan os.Signal),
	}
	s.Server = &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	if err := s.connect(driver, dsn); err != nil {
		return nil, err
	}

	s.rest = handlers.REST{
		Restaurants: mariadb.RestaurantModel{DB: db},
	}
	return &s, nil
}

// connect establishes a database connection using the sqlx library.
func (s *Server) connect(driver, dsn string) error {
	var err error

	db, err = sqlx.Connect(driver, dsn)
	if err != nil {
		return errors.Wrap(err, "connection failed")
	}

	return nil
}

// RunMigration sets up all tables by invoking the individual Migrate() methods.
func (s *Server) RunMigration() error {
	err := s.rest.Restaurants.Migrate()
	return err
}

// Run mounts all API routes, establishes a database connection and starts
// listening to the specified port. The server can be shut down with Ctrl + C.
func (s *Server) Run() error {
	s.mountRoutes()
	signal.Notify(s.interrupt, os.Interrupt)

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	<-s.interrupt
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)

	if err := s.Shutdown(timeout); err != nil {
		return err
	}

	defer cancel()
	db.Close()

	return nil
}

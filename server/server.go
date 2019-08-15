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
	"context"
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/gateways"
	"github.com/dominikbraun/foodunit/gateways/mariadb"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type SessionMode int

const (
	ResumeSessions  SessionMode = 0
	DiscardSessions SessionMode = 1
)

type Server struct {
	*http.Server
	router       *chi.Mux
	conf         *gateways.Conf
	initGateways func() error
	interrupt    chan os.Signal
	sessMode     SessionMode
}

func New(conf *gateways.Conf, sessMode SessionMode) *Server {
	srv := Server{
		router:       createRouter(),
		conf:         conf,
		initGateways: registerGateways,
		interrupt:    make(chan os.Signal),
		sessMode:     sessMode,
	}

	srv.Server = &http.Server{
		Addr:    ":8000",
		Handler: srv.router,
	}

	return &srv
}

func (srv *Server) Start() {
	srv.mountRoutes()
	_ = srv.initGateways()

	if err := mariadb.Connect(srv.conf); err != nil {
		log.Fatalf("could not connect to '%s' as %s\n", srv.conf.DBName, srv.conf.User)
	}

	signal.Notify(srv.interrupt, os.Interrupt)

	go func() {
		<-srv.interrupt
		ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)

		err := srv.Shutdown(ctx)
		if err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
		defer cancelFn()
	}()

	log.Fatal(srv.ListenAndServe())
}

func createRouter() *chi.Mux {
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

func registerGateways() error {
	load.RepositoryLoader.Users = mariadb.UserRepository{}
	load.RepositoryLoader.Suppliers = mariadb.SupplierRepository{}
	load.RepositoryLoader.Categories = mariadb.CategoryRepository{}
	load.RepositoryLoader.Dishes = mariadb.DishRepository{}
	load.RepositoryLoader.Offers = mariadb.OfferRepository{}
	load.RepositoryLoader.Orders = mariadb.OrderRepository{}
	load.RepositoryLoader.Positions = mariadb.PositionRepository{}
	load.RepositoryLoader.IsReady = true

	return nil
}

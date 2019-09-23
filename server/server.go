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
	"context"
	"github.com/dominikbraun/foodunit/services/offer"
	"github.com/dominikbraun/foodunit/services/restaurant"
	"github.com/dominikbraun/foodunit/services/user"
	"github.com/dominikbraun/foodunit/session"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/dominikbraun/foodunit/storage/maria"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	Driver string `json:"driver"`
	DSN    string `json:"dsn"`
}

type Server struct {
	router  *chi.Mux
	http    *http.Server
	db      *sqlx.DB
	session session.Manager

	restaurants     storage.Restaurant
	categories      storage.Category
	dishes          storage.Dish
	characteristics storage.Characteristic
	variants        storage.Variant
	users           storage.User
	offers          storage.Offer
	orders          storage.Order
	positions       storage.Position

	restaurantService *restaurant.Service
	userService       *user.Service
	offerService      *offer.Service
}

func New(config *Config) (*Server, error) {
	s := Server{}
	var err error

	s.router = newRouter()
	s.http = newHTTPServer(s.router)
	s.db, err = sqlx.Open(config.Driver, config.DSN)
	s.session = session.NewManager()

	s.restaurants = maria.NewRestaurant(s.db)
	s.categories = maria.NewCategory(s.db)
	s.dishes = maria.NewDish(s.db)
	s.characteristics = maria.NewCharacteristic(s.db)
	s.variants = maria.NewVariant(s.db)
	s.users = maria.NewUser(s.db)
	s.offers = maria.NewOffer(s.db)
	s.orders = maria.NewOrder(s.db)
	s.positions = maria.NewPosition(s.db)

	s.restaurantService = restaurant.NewService(s.restaurants, s.categories, s.dishes)
	s.userService = user.NewService(s.users)
	s.offerService = offer.NewService(s.offers, s.orders, s.positions)

	s.mountRoutes()

	return &s, err
}

func (s *Server) Run() {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		if err := s.http.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}

func newRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)
	return router
}

func newHTTPServer(handler http.Handler) *http.Server {
	server := http.Server{
		Addr:    "9292",
		Handler: handler,
	}
	return &server
}

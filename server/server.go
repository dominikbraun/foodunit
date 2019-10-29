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
	conf "github.com/dominikbraun/foodunit/config"
	"github.com/dominikbraun/foodunit/controllers/rest"
	"github.com/dominikbraun/foodunit/services/dish"
	"github.com/dominikbraun/foodunit/services/mail"
	"github.com/dominikbraun/foodunit/services/offer"
	"github.com/dominikbraun/foodunit/services/order"
	"github.com/dominikbraun/foodunit/services/restaurant"
	"github.com/dominikbraun/foodunit/services/user"
	"github.com/dominikbraun/foodunit/session"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/dominikbraun/foodunit/storage/maria"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Config concludes important configuration values which will be consumed
// by the API server.
type Config struct {
	Driver string `json:"driver"`
	DSN    string `json:"dsn"`
	Addr   string `json:"addr"`
}

// Server represents an full-featured, executable REST API server. It is the
// central place where all individual components are being wired up.
//
// It initializes and holds all storage objects as well as the business services.
// Each service requires multiple storage objects and the server will pass its
// own objects by reference to the corresponding services.
//
// Since a service assumes that the storage objects are ready to use, each storage
// is initialized first and will then be passed to the service.
type Server struct {
	router  *chi.Mux
	http    *http.Server
	db      *sqlx.DB
	session session.Manager

	appConfig conf.Reader

	restaurants     storage.Restaurant
	categories      storage.Category
	dishes          storage.Dish
	characteristics storage.Characteristic
	variants        storage.Variant
	users           storage.User
	offers          storage.Offer
	orders          storage.Order
	positions       storage.Position
	configurations  storage.Configuration

	mailService *mail.Service

	restaurantService *restaurant.Service
	dishService       *dish.Service
	userService       *user.Service
	offerService      *offer.Service
	orderService      *order.Service

	controller *rest.Controller
}

// New is the actual entry point for the API server binary. It initializes all
// server components like the internal HTTP server or the router. Most importantly,
// it will create a database connection pool which will be passed to the storage
// factories, and the resulting storage object will be passed to the services.
//
// This function is also the place where dependency injection happens: The server
// decides which storage implementations will be used - e. g. maria.Restaurant for
// the Server.restaurant field.
func New(config *Config) (*Server, error) {
	s := Server{}
	var err error

	s.router = newRouter()
	s.http = newHTTPServer(s.router, config.Addr)
	s.db, err = sqlx.Open(config.Driver, config.DSN)
	s.session = session.NewManager(s.db)

	if s.appConfig, err = conf.New("app"); err != nil {
		return nil, err
	}

	s.restaurants = maria.NewRestaurant(s.db)
	s.categories = maria.NewCategory(s.db)
	s.dishes = maria.NewDish(s.db)
	s.characteristics = maria.NewCharacteristic(s.db)
	s.variants = maria.NewVariant(s.db)
	s.users = maria.NewUser(s.db)
	s.offers = maria.NewOffer(s.db)
	s.orders = maria.NewOrder(s.db)
	s.positions = maria.NewPosition(s.db)
	s.configurations = maria.NewConfiguration(s.db)

	sgAPIKey := s.appConfig.GetString("sendgrid_api_key")
	s.mailService = mail.New(sgAPIKey, s.users)

	s.restaurantService = restaurant.NewService(s.restaurants, s.categories, s.dishes)
	s.dishService = dish.NewService(s.dishes, s.characteristics, s.variants)
	s.userService = user.NewService(s.appConfig, s.users, s.mailService)
	s.offerService = offer.NewService(s.appConfig, s.restaurants, s.users, s.offers, s.orders, s.positions, s.mailService)
	s.orderService = order.NewService(s.offers, s.orders, s.positions, s.configurations, s.dishes, s.characteristics, s.variants)

	s.controller = rest.NewController(
		s.restaurantService, s.dishService, s.userService, s.offerService, s.orderService,
	)

	s.mountRoutes()

	return &s, err
}

// Run starts the internal HTTP server which makes the API server listen on the
// port specified in the passed Config instance. This function will panic in case
// an error occurs.
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

// newRouter creates a new router instance which will be used as the server's router.
// The actual API endpoints will be generated and mounted in mountRoutes.
func newRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		render.SetContentType(render.ContentTypeJSON),
	)

	corsSettings := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(corsSettings.Handler)

	return router
}

// newHTTPServer creates a new HTTP server instance which will be used as the
// server's internal HTTP server. It is configured to run on the specified port.
func newHTTPServer(handler http.Handler, addr string) *http.Server {
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return &server
}

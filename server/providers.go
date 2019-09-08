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
	"github.com/dominikbraun/foodunit/controllers"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/dominikbraun/foodunit/storage/mariadb"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// provideDB provides a ready-to-go database connection pool.
func provideDBConnection(driver, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "connection failed")
	}

	return db, err
}

// provideRouter provides a routing service instance.
func provideRouter() *chi.Mux {
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

// provideRestaurantModel provides a storage.RestaurantModel implementation.
func provideRestaurantModel(db *sqlx.DB) storage.RestaurantModel {
	restaurantModel := mariadb.RestaurantModel{
		DB: db,
	}
	return &restaurantModel
}

// provideRestaurantModel provides a storage.RestaurantModel implementation.
func provideUserModel(db *sqlx.DB) storage.UserModel {
	userModel := mariadb.UserModel{
		DB: db,
	}
	return &userModel
}

// provideRESTController provides a controller instance for handing REST requests.
func provideRESTController(r storage.RestaurantModel, u storage.UserModel) *controllers.REST {
	controller := controllers.REST{
		Restaurants: r,
		Users:       u,
	}
	return &controller
}

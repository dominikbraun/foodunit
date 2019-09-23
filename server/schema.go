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
	"github.com/dominikbraun/foodunit/storage"
	"log"
)

func (s *Server) RunSchemaSetup(drop bool) {
	entities := []storage.Entity{
		s.restaurants,
		s.categories,
		s.dishes,
		s.characteristics,
		s.variants,
		s.users,
		s.offers,
		s.orders,
		s.positions,
	}
	if err := CreateSchema(entities, drop); err != nil {
		log.Fatal()
	}
}

func CreateSchema(entities []storage.Entity, drop bool) error {
	var err error

	if drop {
		for _, s := range entities {
			err = s.Drop()
			if err != nil {
				return err
			}
		}
	}

	for _, s := range entities {
		err = s.Create()
		if err != nil {
			return err
		}
	}

	return nil
}

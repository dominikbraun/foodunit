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

// Package dishes provides services and types for Dish-related data.
package dishes

import (
	"github.com/dominikbraun/foodunit/storage"
)

type Service struct {
	dishes          storage.Dish
	characteristics storage.Characteristic
	variants        storage.Variant
}

func NewService(d storage.Dish, c storage.Characteristic, v storage.Variant) *Service {
	service := Service{
		dishes:          d,
		characteristics: c,
		variants:        v,
	}
	return &service
}

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

// Package load provides concrete interface implementations prescribed
// by the Domain Language.
package load

import "github.com/dominikbraun/foodunit/dl"

// RepositoryLoader is the global instance which will be accessed by the
// interactors after it has been initialized by the corresponding adapter.
var RepositoryLoader *repositoryLoader

// repositoryLoader holds functions providing repository implementations.
type repositoryLoader struct {
	IsReady    bool
	Users      dl.UserRepository
	Suppliers  dl.SupplierRepository
	Categories dl.CategoryRepository
	Dishes     dl.DishRepository
	Offers     dl.OfferRepository
	Orders     dl.OrderRepository
	Positions  dl.PositionRepository
}

// init initializes the global repositoryLoader instance so that it
// can be used by an interface adapter out of the box.
func init() {
	RepositoryLoader = &repositoryLoader{}
	RepositoryLoader.IsReady = false
}

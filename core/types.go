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

// Package core provides business logic interactors and services.
package core

// Restaurant info represents a simple conclusion of relevant meta information.
type RestaurantInfo struct {
	Name       string `db:"name" json:"name"`
	Street     string `db:"street" json:"street"`
	PostalCode string `db:"postal_code" json:"postal_code"`
	City       string `db:"city" json:"city"`
	Phone      string `db:"phone" json:"phone"`
	Open       string `db:"open" json:"open"`
	Website    string `db:"website" json:"website"`
}

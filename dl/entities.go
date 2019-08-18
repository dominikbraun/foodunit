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

// Package dl provides domain language entities and rules.

package dl

import "time"

// Supplier represents a food supplier like a restaurant or delivery service.
type Supplier struct {
	ID         uint64 `db:"id"`
	Name       string `db:"name"`
	Street     string `db:"street"`
	PostalCode string `db:"postal_code"`
	City       string `db:"city"`
	OpenMon    string `db:"open_mon"`
	OpenTue    string `db:"open_tue"`
	OpenWed    string `db:"open_wed"`
	OpenThu    string `db:"open_thu"`
	OpenFri    string `db:"open_fri"`
	OpenSat    string `db:"open_sat"`
	OpenSun    string `db:"open_sun"`
	Website    string `db:"website"`
}

// Category represents the category or menu's section a dish belongs to.
type Category struct {
	ID         uint64 `db:"id"`
	SupplierID uint64 `db:"supplier_id"`
	Name       string `db:"name"`
	ImgPath    string `db:"img_path"`
}

// Dish represents a meal that is offered by a Supplier.
type Dish struct {
	ID          uint64 `db:"id"`
	CategoryID  uint64 `db:"category_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       uint8  `db:"price"`
}

// User represents a person that creates offers and orders food.
type User struct {
	ID   uint64 `db:"id"`
	Mail string `db:"mail"`
	Name string `db:"name"`
}

// Offer represents an User's offer to order food for their team.
type Offer struct {
	ID         uint64    `db:"id"`
	UserID     uint64    `db:"user_id"`
	SupplierID uint64    `db:"supplier_id"`
	ValidFrom  time.Time `db:"valid_from"`
	ValidTo    time.Time `db:"valid_to"`
	IsPlaced   bool      `db:"is_placed"`
	PickupInfo string    `db:"pickup_info"`
}

// Order represents an User's order that was placed as part of
// an Offer. The list of ordered food is depicted in Positions.
type Order struct {
	ID      uint64 `db:"id"`
	UserID  uint64 `db:"user_id"`
	OfferID uint64 `db:"offer_id"`
}

// Position represents one of multiple items contained in an Order.
type Position struct {
	ID      uint64 `db:"id"`
	OrderID uint64 `db:"order_id"`
	DishID  uint64 `db:"dish_id"`
	Note    string `db:"note"`
}

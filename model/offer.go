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

// Package model provides basic domain models.
package model

import "time"

// Offer represents an user's offer for ordering food for their team at
// a restaurant. If the food won't be delivered by the restaurant/delivery
// service, the owner of the offer is responsible for picking up the ordered food.
type Offer struct {
	ID            uint64     `db:"id"`
	Owner         User       `db:"owner_user_id"`
	Restaurant    Restaurant `db:"restaurant_id"`
	ValidFrom     time.Time  `db:"valid_from"`
	ValidTo       time.Time  `db:"valid_to"`
	Responsible   User       `db:"responsible_user_id"`
	IsPlaced      bool       `db:"is_placed"`
	IsCancelled   bool       `db:"is_cancelled"`
	ReadyAt       time.Time  `db:"ready_at"`
	PaypalEnabled bool       `db:"paypal_enabled"`
	Orders        []Order    `db:"orders"`
}

// Order represents an user's order for a particular offer. In most cases, there
// will be multiple orders (i. e. multiple users ordering) at an offer.
type Order struct {
	ID        uint64     `db:"id"`
	User      User       `db:"user_id"`
	Positions []Position `db:"positions"`
	IsPaid    bool       `db:"is_paid"`
}

// Position represents an order's position. It includes the actual dish ordered,
// next to an alternative dish if the actual dish isn't available for sure.
// A position also hold several configurations.
type Position struct {
	ID             uint64          `db:"id"`
	Dish           Dish            `db:"dish_id"`
	Alternative    Dish            `db:"alternative_dish_id"`
	Note           string          `db:"note"`
	Configurations []Configuration `db:"configurations"`
}

// Configuration represents the actual choice a user has made for a characteristic. If the
// user orders a pizza and chooses `pizza size: 30 cm`, the order position includes
// a Configuration saying that `pizza size` = `30 cm`. Therefore, it's simply a mapping
// of a dish's characteristic against one ore more concrete variants.
type Configuration struct {
	ID             uint64         `db:"id"`
	Characteristic Characteristic `db:"characteristic_id"`
	Variants       []Variant      `db:"variants"`
}

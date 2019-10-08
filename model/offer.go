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

type Order struct {
	ID        uint64     `db:"id"`
	User      User       `db:"user_id"`
	Positions []Position `db:"positions"`
	IsPaid    bool       `db:"is_paid"`
}

type Position struct {
	ID             uint64          `db:"id"`
	Dish           Dish            `db:"dish_id"`
	Alternative    Dish            `db:"alternative_dish_id"`
	Note           string          `db:"note"`
	Configurations []Configuration `db:"configurations"`
}

type Configuration struct {
	ID             uint64         `db:"id"`
	Characteristic Characteristic `db:"characteristic_id"`
	Variants       []Variant      `db:"variants"`
}

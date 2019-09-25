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

// Package offer provides services and types for Offer-related data.
package offer

import "time"

type Creation struct {
	Restaurant    uint64    `json:"restaurant_id"`
	ValidFrom     time.Time `json:"valid_from"`
	ValidTo       time.Time `json:"valid_to"`
	PaypalEnabled bool      `json:"paypal_enabled"`
}

type ActiveOffer struct {
	ID            uint64     `json:"id"`
	Owner         User       `json:"owner"`
	Restaurant    Restaurant `json:"restaurant"`
	ValidFrom     time.Time  `json:"valid_from"`
	ValidTo       time.Time  `json:"valid_to"`
	PaypalEnabled bool       `json:"paypal_enabled"`
}

type User struct {
	Name string `json:"name"`
}

type Restaurant struct {
	Name string `json:"name"`
}

type View struct {
	ID            uint64    `db:"id"`
	Owner         User      `db:"owner"`
	ValidFrom     time.Time `db:"valid_from"`
	ValidTo       time.Time `db:"valid_to"`
	Responsible   User      `db:"responsible_user_id"`
	IsPlaced      bool      `db:"is_placed"`
	ReadyAt       time.Time `db:"ready_at"`
	PaypalEnabled bool      `db:"paypal_enabled"`
}

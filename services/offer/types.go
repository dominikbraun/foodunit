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

// Creation represents the creation of a new offer.
type Creation struct {
	Restaurant    uint64    `json:"restaurant_id"`
	ValidFrom     time.Time `json:"valid_from"`
	ValidTo       time.Time `json:"valid_to"`
	PaypalEnabled bool      `json:"paypal_enabled"`
}

// Offer is the API output for model.Offer.
type Offer struct {
	ID            uint64     `json:"id"`
	Owner         User       `json:"owner"`
	Restaurant    Restaurant `json:"restaurant"`
	ValidFrom     time.Time  `json:"valid_from"`
	ValidTo       time.Time  `json:"valid_to"`
	PaypalEnabled bool       `json:"paypal_enabled"`
}

// User is the API output for model.User.
type User struct {
	ID uint64 `json:"id"`
}

// Restaurant is the API output for model.Restaurant.
type Restaurant struct {
	ID uint64 `json:"id"`
}

// View represents an offer view.
type View struct {
	ID            uint64    `json:"id"`
	Owner         User      `json:"owner"`
	ValidFrom     time.Time `json:"valid_from"`
	ValidTo       time.Time `json:"valid_to"`
	Responsible   User      `json:"responsible_user_id"`
	IsPlaced      bool      `json:"is_placed"`
	IsCancelled   bool      `json:"is_cancelled"`
	ReadyAt       time.Time `json:"ready_at"`
	PaypalEnabled bool      `json:"paypal_enabled"`
}

// ReadAtSetter represents a setter type for the ReadyAt property.
type ReadyAtSetter struct {
	ReadyAt time.Time `json:"ready_at"`
}

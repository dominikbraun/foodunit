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

// Package dto provides data transfer objects as input and output data.
package dto

import "time"

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

// UserRegistration concludes required data for registering a new user.
type UserRegistration struct {
	MailAddr       string `db:"mail_addr" json:"mail_addr"`
	Name           string `db:"name" json:"name"`
	PaypalMailAddr string `db:"paypal_mail_addr" json:"paypal_mail_addr"`
	Password       string `db:"password" json:"password"`
}

// UserLogin provides data required for logging in.
type UserLogin struct {
	MailAddr string `db:"mail_addr" json:"mail_addr"`
	Password string `db:"password" json:"password"`
}

// NewOffer provides data required for creating an offer in.
type NewOffer struct {
	Restaurant    uint64    `json:"restaurant"`
	ValidFrom     time.Time `json:"valid_from"`
	ValidTo       time.Time `json:"valid_to"`
	Responsible   uint64    `json:"responsible"`
	IsPlaced      bool      `json:"is_placed"`
	ReadyAt       time.Time `json:"ready_at"`
	PaypalEnabled bool      `json:"paypal_enabled"`
}

// Offers provides a list of offer ids
type Offers struct {
	Offers []uint64 `json:"offer_ids"`
}

// Menu represents a restaurant's menu. It consists of multiple categories.
type Menu struct {
	Categories []MenuCategory `json:"categories"`
}

// MenuCategory represents a section in a menu, holding several dishes.
type MenuCategory struct {
	Name   string     `json:"name"`
	Dishes []MenuDish `json:"dishes"`
}

// MenuDish is a dish as it appears in a restaurant's menu.
type MenuDish struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        uint   `json:"price"`
	IsUncertain  bool   `json:"is_uncertain"`
	IsHealthy    bool   `json:"is_healthy"`
	IsVegetarian bool   `json:"is_vegetarian"`
}

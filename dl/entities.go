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

// Entity represents any domain entity which will be persisted.
type Entity interface {
	GetID() uint64
}

// Model represents an Entity implementation. It can be embedded in any
// struct so that it becomes a storable entity.
type Model struct {
	ID uint64 `db:"id" json:"id"`
}

// GetID implements Entity.GetID.
func (m Model) GetID() uint64 {
	return m.ID
}

// Restaurant represents any type of restaurant or delivery service.
type Restaurant struct {
	Model
	Name       string     `db:"name" json:"name"`
	Street     string     `db:"street" json:"street"`
	PostalCode string     `db:"postal_code" json:"postal_code"`
	City       string     `db:"city" json:"city"`
	Phone      string     `db:"phone" json:"phone"`
	OpenMon    string     `db:"open_mon" json:"open_mon"`
	OpenTue    string     `db:"open_tue" json:"open_tue"`
	OpenWed    string     `db:"open_wed" json:"open_wed"`
	OpenThu    string     `db:"open_thu" json:"open_thu"`
	OpenFri    string     `db:"open_fri" json:"open_fri"`
	OpenSat    string     `db:"open_sat" json:"open_sat"`
	OpenSun    string     `db:"open_sun" json:"open_sun"`
	Website    string     `db:"website" json:"website"`
	IsActive   bool       `db:"is_active" json:"is_active"`
	Menu       []Category `db:"menu" json:"menu"`
}

// Category represents a menu section holding multiple dishes.
type Category struct {
	Model
	Name   string `db:"name" json:"name"`
	Dishes []Dish `db:"dishes" json:"dishes"`
}

// Dish represents a meal that is offered by a restaurant.
type Dish struct {
	Model
	Name            string           `db:"name" json:"name"`
	Description     string           `db:"description" json:"description"`
	Price           uint             `db:"price" json:"price"`
	IsUncertain     bool             `db:"is_uncertain" json:"is_uncertain"`
	IsHealthy       bool             `db:"is_healthy" json:"is_healthy"`
	Characteristics []Characteristic `db:"characteristics" json:"characteristics"`
}

// Characteristic represents a user-configurable dish feature.
type Characteristic struct {
	Model
	Name     string    `db:"name" json:"name"`
	Multiple bool      `db:"multiple" json:"multiple"`
	Variants []Variant `db:"variants" json:"variants"`
}

// Variant represents a discrete value of a characteristic.
type Variant struct {
	Model
	Value     string `db:"value" json:"value"`
	IsDefault bool   `db:"is_default" json:"is_default"`
	Price     uint   `db:"price" json:"price"`
}

// User represents a person that creates offers and orders food.
type User struct {
	Model
	MailAddr       string    `db:"mail_addr" json:"mail_addr"`
	Name           string    `db:"name" json:"name"`
	IsAdmin        bool      `db:"is_admin" json:"is_admin"`
	PaypalMailAddr string    `db:"paypal_mail_addr" json:"paypal_mail_addr"`
	Score          int       `db:"score" json:"score"`
	PasswordHash   []byte    `db:"password_hash" json:"password_hash"`
	Created        time.Time `db:"created" json:"created"`
}

// Offer represents an user's offer to order food for their friends or team.
type Offer struct {
	Model
	Owner         User       `db:"owner_user_id" json:"owner_user_id"`
	Restaurant    Restaurant `db:"restaurant" json:"restaurant"`
	ValidFrom     time.Time  `db:"valid_from" json:"valid_from"`
	ValidTo       time.Time  `db:"valid_to" json:"valid_to"`
	Responsible   User       `db:"responsible_user_id" json:"responsible_user_id"`
	IsPlaced      bool       `db:"is_placed" json:"is_placed"`
	ReadyAt       time.Time  `db:"ready_at" json:"ready_at"`
	PaypalEnabled bool       `db:"paypal_enabled" json:"paypal_enabled"`
	Orders        []Order    `db:"orders" json:"orders"`
}

// Order represents an user's order that was placed as part of an order.
type Order struct {
	Model
	User      User       `db:"user_id" json:"user_id"`
	Positions []Position `db:"positions" json:"positions"`
	IsPaid    bool       `db:"is_paid" json:"is_paid"`
}

// Position represents one of multiple dishes contained in an user order.
type Position struct {
	Model
	Dish        Dish   `db:"dish_id" json:"dish_id"`
	Alternative Dish   `db:"alternative" json:"alternative"`
	Note        string `db:"note" json:"note"`
}

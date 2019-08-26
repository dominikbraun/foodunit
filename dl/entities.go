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

// Entity represents any domain entity which will be stored in a repository.
type Entity interface {
	GetID() uint64
}

// Model represents a discrete Entity implementation. It will be implemented
// by all actual entities so that they get an ID field and become storable.
type Model struct {
	ID uint64 `db:"id" json:"id"`
}

// GetID returns a model's ID.
func (e Model) GetID() uint64 {
	return e.ID
}

// Supplier represents a food supplier like a restaurant or delivery service.
type Supplier struct {
	Model
	Name       string `db:"name" json:"name"`
	Street     string `db:"street" json:"street"`
	PostalCode string `db:"postal_code" json:"postal_code"`
	City       string `db:"city" json:"city"`
	Phone      string `db:"phone" json:"phone"`
	OpenMon    string `db:"open_mon" json:"open_mon"`
	OpenTue    string `db:"open_tue" json:"open_tue"`
	OpenWed    string `db:"open_wed" json:"open_wed"`
	OpenThu    string `db:"open_thu" json:"open_thu"`
	OpenFri    string `db:"open_fri" json:"open_fri"`
	OpenSat    string `db:"open_sat" json:"open_sat"`
	OpenSun    string `db:"open_sun" json:"open_sun"`
	Website    string `db:"website" json:"website"`
	IsActive   bool   `db:"is_active" json:"is_active"`
}

// Category represents the category or menu section a dish belongs to.
type Category struct {
	Model
	SupplierID uint64 `db:"supplier_id" json:"supplier_id"`
	Name       string `db:"name" json:"name"`
	ImgPath    string `db:"img_path" json:"img_path"`
}

// Dish represents a meal that is offered by a Supplier.
type Dish struct {
	Model
	CategoryID  uint64 `db:"category_id" json:"category_id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Price       uint8  `db:"price" json:"price"`
	IsUncertain bool   `db:"is_uncertain" json:"is_uncertain"`
}

// Characteristic represents a user-configurable product feature.
type Characteristic struct {
	Model
	Name     string `db:"name" json:"name"`
	Multiple bool   `db:"multiple" json:"multiple"`
}

// Variant represents a discrete value of a characteristic.
type Variant struct {
	Model
	CharacteristicID uint64 `db:"characteristic_id" json:"characteristic_id"`
	Value            string `db:"value" json:"value"`
}

// User represents a person that creates offers and orders food.
type User struct {
	Model
	Mail       string `db:"mail" json:"mail"`
	Name       string `db:"name" json:"name"`
	PayPalMail string `db:"pay_pal_mail" json:"pay_pal_mail"`
}

// Offer represents an User's offer to order food for their team.
type Offer struct {
	Model
	UserID     uint64    `db:"user_id" json:"user_id"`
	SupplierID uint64    `db:"supplier_id" json:"supplier_id"`
	ValidFrom  time.Time `db:"valid_from" json:"valid_from"`
	ValidTo    time.Time `db:"valid_to" json:"valid_to"`
	IsPlaced   bool      `db:"is_placed" json:"is_placed"`
	PickupInfo string    `db:"pickup_info" json:"pickup_info"`
}

// Order represents an User's order that was placed as part of an Offer. The
// list of ordered food is depicted in Positions.
type Order struct {
	Model
	UserID  uint64 `db:"user_id" json:"user_id"`
	OfferID uint64 `db:"offer_id" json:"offer_id"`
	IsPaid  bool   `db:"is_paid" json:"is_paid"`
}

// Position represents one of multiple items contained in an Order.
type Position struct {
	Model
	OrderID uint64 `db:"order_id" json:"order_id"`
	DishID  uint64 `db:"dish_id" json:"dish_id"`
	Note    string `db:"note" json:"note"`
}

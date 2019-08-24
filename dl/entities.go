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
	ID uint64 `db:"id"`
}

// GetID returns a model's ID.
func (e Model) GetID() uint64 {
	return e.ID
}

// Supplier represents a food supplier like a restaurant or delivery service.
type Supplier struct {
	Model
	Name       string `db:"name"`
	Street     string `db:"street"`
	PostalCode string `db:"postal_code"`
	City       string `db:"city"`
	Phone      string `db:"phone"`
	OpenMon    string `db:"open_mon"`
	OpenTue    string `db:"open_tue"`
	OpenWed    string `db:"open_wed"`
	OpenThu    string `db:"open_thu"`
	OpenFri    string `db:"open_fri"`
	OpenSat    string `db:"open_sat"`
	OpenSun    string `db:"open_sun"`
	Website    string `db:"website"`
	IsActive   bool   `db:"is_active"`
}

// Category represents the category or menu section a dish belongs to.
type Category struct {
	Model
	SupplierID uint64 `db:"supplier_id"`
	Name       string `db:"name"`
	ImgPath    string `db:"img_path"`
}

// Dish represents a meal that is offered by a Supplier.
type Dish struct {
	Model
	CategoryID  uint64 `db:"category_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       uint8  `db:"price"`
}

// Characteristic represents a user-configurable product feature.
type Characteristic struct {
	Model
	Name     string `db:"name"`
	Multiple bool   `db:"multiple"`
}

// Variant represents a discrete value of a characteristic.
type Variant struct {
	Model
	CharacteristicID uint64 `db:"characteristic_id"`
	Value            string `db:"value"`
}

// User represents a person that creates offers and orders food.
type User struct {
	Model
	Mail       string `db:"mail"`
	Name       string `db:"name"`
	PayPalMail string `db:"pay_pal_mail"`
}

// Offer represents an User's offer to order food for their team.
type Offer struct {
	Model
	UserID     uint64    `db:"user_id"`
	SupplierID uint64    `db:"supplier_id"`
	ValidFrom  time.Time `db:"valid_from"`
	ValidTo    time.Time `db:"valid_to"`
	IsPlaced   bool      `db:"is_placed"`
	PickupInfo string    `db:"pickup_info"`
}

// Order represents an User's order that was placed as part of an Offer. The
// list of ordered food is depicted in Positions.
type Order struct {
	Model
	UserID  uint64 `db:"user_id"`
	OfferID uint64 `db:"offer_id"`
	IsPaid  bool   `db:"is_paid"`
}

// Position represents one of multiple items contained in an Order.
type Position struct {
	Model
	OrderID uint64 `db:"order_id"`
	DishID  uint64 `db:"dish_id"`
	Note    string `db:"note"`
}

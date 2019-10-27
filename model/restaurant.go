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

// Restaurant represents a restaurant or delivery service. When a user creates
// an offer, the offer will be associated with exactly one Restaurant.
type Restaurant struct {
	ID         uint64     `db:"id"`
	Name       string     `db:"name"`
	Street     string     `db:"street"`
	PostalCode string     `db:"postal_code"`
	City       string     `db:"city"`
	Phone      string     `db:"phone"`
	OpenMon    string     `db:"open_mon"`
	OpenTue    string     `db:"open_tue"`
	OpenWed    string     `db:"open_wed"`
	OpenThu    string     `db:"open_thu"`
	OpenFri    string     `db:"open_fri"`
	OpenSat    string     `db:"open_sat"`
	OpenSun    string     `db:"open_sun"`
	Website    string     `db:"website"`
	IsActive   bool       `db:"is_active"`
	Menu       []Category `db:"menu"`
}

// Category represents a menu section or dish category. Each category is
// associated with only one Restaurant.
type Category struct {
	ID     uint64 `db:"id"`
	Name   string `db:"name"`
	Dishes []Dish `db:"dishes"`
}

// Dish is a dish or meal offered by a restaurant. Each dish is associated
// with only one Category.
type Dish struct {
	ID              uint64           `db:"id"`
	Name            string           `db:"name"`
	Description     string           `db:"description"`
	Price           uint             `db:"price"`
	IsUncertain     bool             `db:"is_uncertain"`
	IsHealthy       bool             `db:"is_healthy"`
	IsVegetarian    bool             `db:"is_vegetarian"`
	Characteristics []Characteristic `db:"characteristics"`
}

// Characteristic represents a property of a dish that can be configured individually
// by an user, for example `pizza size`. Each dish can have multiple characteristics,
// and each characteristic may be associated with multiple dishes (since `pizza size`
// will exist for both `Pizza Salami` and `Pizza Margherita`).
type Characteristic struct {
	ID       uint64    `db:"id"`
	Name     string    `db:"name"`
	Multiple bool      `db:"multiple"`
	Variants []Variant `db:"variants"`
}

// Variant represents a concrete value for a characteristic, for example `28 cm` for
// `pizza size`. It may be used for extra ingredients and therefore has a price.
type Variant struct {
	ID        uint64 `db:"id"`
	Value     string `db:"value"`
	IsDefault bool   `db:"is_default"`
	Price     uint   `db:"price"`
}

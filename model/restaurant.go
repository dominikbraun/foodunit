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

type Restaurant struct {
	ID         uint64     `db:"id" json:"id"`
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

type Category struct {
	ID     uint64 `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Dishes []Dish `db:"dishes" json:"dishes"`
}

type Dish struct {
	ID              uint64           `db:"id" json:"id"`
	Name            string           `db:"name" json:"name"`
	Description     string           `db:"description" json:"description"`
	Price           uint             `db:"price" json:"price"`
	IsUncertain     bool             `db:"is_uncertain" json:"is_uncertain"`
	IsHealthy       bool             `db:"is_healthy" json:"is_healthy"`
	IsVegetarian    bool             `db:"is_vegetarian" json:"is_vegetarian"`
	Characteristics []Characteristic `db:"characteristics" json:"characteristics"`
}

type Characteristic struct {
	ID       uint64    `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Multiple bool      `db:"multiple" json:"multiple"`
	Variants []Variant `db:"variants" json:"variants"`
}

type Variant struct {
	ID        uint64 `db:"id" json:"id"`
	Value     string `db:"value" json:"value"`
	IsDefault bool   `db:"is_default" json:"is_default"`
	Price     uint   `db:"price" json:"price"`
}

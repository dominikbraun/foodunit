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

// Package order provides services and types for Order-related data.
package order

type Order struct {
	ID        uint64     `json:"id"`
	User      User       `json:"user"`
	Positions []Position `json:"positions"`
	IsPaid    bool       `json:"is_paid"`
	Total     uint       `json:"total"`
}

type User struct {
	Name string `json:"name"`
}

type Position struct {
	ID             uint64          `json:"id"`
	Dish           Dish            `json:"dish"`
	Alternative    Dish            `json:"alternative_dish"`
	Note           string          `json:"note"`
	Configurations []Configuration `json:"configurations"`
}

type Dish struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type Configuration struct {
	CharacteristicName string    `json:"characteristic_name"`
	Multiple           bool      `json:"multiple"`
	Variants           []Variant `json:"variants"`
}

type Variant struct {
	ID    uint64 `json:"id"`
	Value string `json:"value"`
	Price uint   `json:"price"`
}

type Update struct {
	OfferID   uint64           `json:"offer_id"`
	UserID    uint64           `json:"user_id"`
	Positions []UpdatePosition `json:"positions"`
}

type UpdatePosition struct {
	DishID            uint64                `json:"dish_id"`
	AlternativeDishID uint64                `json:"alternative_dish_id"`
	Configurations    []UpdateConfiguration `json:"configurations"`
	Note              string                `json:"note"`
}

type UpdateConfiguration struct {
	CharacteristicID uint64   `json:"characteristic_id"`
	VariantIDs       []uint64 `json:"variants"`
}

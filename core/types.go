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

// Package core provides business services and use case methods.
package core

import "github.com/dominikbraun/foodunit/dl"

// Supplier merely holds a DL Supplier instance.
type Supplier struct {
	dl.Supplier
}

// Menu represents a collection of Categories.
type Menu []Category

// Category represents a menu section listing multiple Dishes.
type Category struct {
	Name    string `json:"name"`
	ImgPath string `json:"img_path"`
	Dishes  []Dish `json:"dishes"`
}

// Dish represents any type of food that will be displayed on the menu.
type Dish struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint8  `json:"price"`
	IsUncertain bool   `json:"is_uncertain"`
}

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

// SupplierService provides methods for Supplier-related use cases.
type SupplierService interface {
	GetInfo(id uint64) Supplier
	//GetMenu(id uint64) []Category
}

// DishService provides methods for Dish-related use cases.
type DishService interface {
	//GetCharacteristics(id uint64) []Characteristic
}

// UserService provides methods for User-related use cases.
type UserService interface {
	//Login(user, pass string) error
	//Register(email, name, pass string) error
	//GetInfo(id uint64) User
}

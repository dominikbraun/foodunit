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

// SupplierRepository provides methods for managing Supplier-related data.
type SupplierRepository interface {
	Create(Supplier) error
	Find(id uint64) (Supplier, error)
	FindBy(field, val string) ([]Supplier, error)
	Update(Supplier) error
}

// DishRepository provides methods for managing Dish-related data.
type DishRepository interface {
	Create(Dish) error
	FindBy(field, val string) ([]Dish, error)
	Update(Dish) error
}

// CharacteristicRepository provides methods for managing Characteristic-related data.
type CharacteristicRepository interface {
	Create(Characteristic) error
	FindBy(field, val string) ([]Characteristic, error)
}

// UserRepository provides methods for managing User-related data.
type UserRepository interface {
	Create(User) error
	Find(id uint64) (User, error)
	Update(User) error
}

// OfferRepository provides methods for managing Offer-related data.
type OfferRepository interface {
	Create(Offer) error
	Find(id uint64) (Offer, error)
	Update(Offer) error
}

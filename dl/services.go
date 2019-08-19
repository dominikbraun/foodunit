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

// SupplierService provides use cases that are implemented in the core package.
type SupplierService interface {
	Get(id uint64) (Supplier, error)
}

// SupplierService provides use cases that are implemented in the core package.
type DishService interface{}

// SupplierService provides use cases that are implemented in the core package.
type UserService interface {
	// Get returns the user which is associated with the session.
	// Get() (User, error)
}

// SupplierService provides use cases that are implemented in the core package.
type OfferService interface {
	// Get(id uint64)
	// GetActive returns all currently active offers.
	// GetActive() ([]Offer, error)
}

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

// Package dl provides Domain Language entities and rules.
package dl

// DataAccess currently provides all necessary repository interfaces as
// methods. Multiple concrete repository implementations can be grouped
// together in an own type - implementing DataAccess - which can be injected
// as a single dependency, instead of multiple separate implementations as
// multiple dependencies.
type DataAccess interface {
	Offers() OfferRepository
	Orders() OrderRepository
	Positions() PositionRepository
	Suppliers() SupplierRepository
	Categories() CategoryRepository
	Dishes() DishRepository
	Users() UserRepository
}

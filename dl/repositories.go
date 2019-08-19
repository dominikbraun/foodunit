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

// DataAccess provides all repository interfaces. A type that implements this
// interface serves as a wrapper for all repository implementations, and can
// therefore be required and injected as a single dependency.
type DataAccess interface {
	Offers() OfferRepository
	Orders() OrderRepository
	Positions() PositionRepository
	Suppliers() SupplierRepository
	Categories() CategoryRepository
	Dishes() DishRepository
	Users() UserRepository
}

// SupplierRepository provides CRUD methods used by the corresponding service.
type SupplierRepository interface {
	Migrate() error
	Create(s *Supplier) (uint64, error)
	Find(id uint64) (*Supplier, error)
	Update(s *Supplier) error
	Delete(s *Supplier) error
}

// CategoryRepository provides CRUD methods used by the corresponding service.
type CategoryRepository interface {
	Migrate() error
	Create(c *Category) (uint64, error)
	Find(id uint64) (*Category, error)
	FindBySupplierID(supplierID uint64) ([]*Category, error)
	Update(c *Category) error
	Delete(c *Category) error
}

// DishRepository provides CRUD methods used by the corresponding service.
type DishRepository interface {
	Migrate() error
	Create(d *Dish) (uint64, error)
	Find(id uint64) (*Dish, error)
	FindByCategoryID(categoryID uint64) ([]*Dish, error)
	Update(d *Dish) error
	Delete(d *Dish) error
}

// UserRepository provides CRUD methods used by the corresponding service.
type UserRepository interface {
	Migrate() error
	Create(u *User) (uint64, error)
	Find(id uint) (*User, error)
	FindByMail(mail string) (*User, error)
	Update(u *User) error
	Delete(u *User) error
}

// OfferRepository provides CRUD methods used by the corresponding service.
type OfferRepository interface {
	Migrate() error
	Create(o *Offer) (uint64, error)
	Find(id uint64) (*Offer, error)
	FindAllActive() ([]*Offer, error)
	Update(o *Offer) error
	Delete(o *Offer) error
}

// OrderRepository provides CRUD methods used by the corresponding service.
type OrderRepository interface {
	Migrate() error
	Create(o *Order) (uint64, error)
	Find(id uint64) (*Order, error)
	FindByOfferID(offerID uint64) ([]*Order, error)
	Update(o *Order) error
	Delete(o *Order) error
}

// PositionRepository provides CRUD methods used by the corresponding service.
type PositionRepository interface {
	Migrate() error
	Create(p *Position) (uint64, error)
	Find(id uint64) (*Position, error)
	FindByOrderID(orderID uint64) ([]*Position, error)
	Update(p *Position) error
	Delete(p *Position) error
}

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

// Item represents a repository item any CRUD operations can be applied to.
type Item interface {
	GetID() string
}

// Repository represents a type of data storage which holds several Items
// and provides methods to CRUD them. A typical example is a MySQL database.
type Repository interface {
	// Store stores an item into the respective table or bucket.
	Store(Item) error
	// Find finds the item with the corresponding ID.
	Find(id uint64) (Item, error)
	// FindBy reads all items that match the condition field = val.
	FindBy(field, val string) ([]Item, error)
	// Update updates the item with the corresponding ID.
	Update(Item) error
	// Delete deletes the item with the corresponding ID.
	Delete(Item) error
}

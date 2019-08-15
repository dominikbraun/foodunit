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

// User represents a person which can login, create offers and order food.
type User struct {
	ID   uint64 `db:"id"`
	Mail string `db:"mail"`
	Name string `db:"name"`
}

// UserRepository provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type UserRepository interface {
	Migrate() error
	Create(u *User) (uint64, error)
	Find(id uint) (*User, error)
	FindByMail(mail string) (*User, error)
	Update(u *User) error
	Delete(u *User) error
}

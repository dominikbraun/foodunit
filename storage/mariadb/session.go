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

// Package mariadb provides MariaDB-compatible model implementations.
package mariadb

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// SessionStorage is a storage.PositionModel implementation.
type SessionStorage struct {
	DB *sqlx.DB
}

// Delete implements storage.Session.Delete.
func (s SessionStorage) Delete(token string) (err error) {
	panic("implement me")
}

// Find implements storage.Session.Find.
func (s SessionStorage) Find(token string) (b []byte, found bool, err error) {
	panic("implement me")
}

// Commit implements storage.Session.Commit.
func (s SessionStorage) Commit(token string, b []byte, expiry time.Time) (err error) {
	panic("implement me")
}

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

// Package session provides interfaces and implementations for session managers.
package session

import "github.com/jmoiron/sqlx"

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	storage := Storage{
		DB: db,
	}
	return &storage
}

func (s *Storage) Create() error {
	query := `
CREATE TABLE sessions (
	token CHAR(43) PRIMARY KEY,
	data BLOB NOT NULL,
	expiry TIMESTAMP(6) NOT NULL
);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *Storage) Drop() error {
	query := `DROP TABLE IF EXISTS sessions`
	_, err := s.DB.Exec(query)

	return err
}

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
	"github.com/dominikbraun/foodunit/dl"
	. "github.com/dominikbraun/foodunit/storage/mariadb/query"
	"github.com/jmoiron/sqlx"
)

// UserModel is a storage.UserModel implementation.
type UserModel struct {
	DB *sqlx.DB
}

// Migrate implements storage.UserModel.Migrate.
func (u UserModel) Migrate() error {
	query := Create("users", Fields{
		"id":               "BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY",
		"mail_addr":        "VARCHAR(254) NOT NULL",
		"name":             "VARCHAR(50) NOT NULL",
		"is_admin":         "BOOLEAN NOT NULL",
		"paypal_mail_addr": "VARCHAR(254) NOT NULL",
		"score":            "INTEGER UNSIGNED NOT NULL",
		"password_hash":    "CHAR(60) NOT NULL",
		"created":          "DATETIME NOT NULL",
	})

	_ = u.DB.MustExec(query)
	return nil
}

// Create implements storage.UserModel.Create.
func (u UserModel) Create(user dl.User) error {
	query := Insert("users", Fields{
		"mail_addr":        "?",
		"name":             "?",
		"is_admin":         "?",
		"paypal_mail_addr": "?",
		"score":            "",
		"password_hash":    "?",
		"created":          "?",
	})

	_, err := u.DB.Exec(query, user.MailAddr, user.Name, user.IsAdmin, user.PaypalMailAddr, user.Score, user.PasswordHash, user.Created)
	if err != nil {
		return err
	}

	return nil
}

// Authenticate implements storage.UserModel.Authenticate.
func (u UserModel) Authenticate(mailAddr, password string) error {
	panic("implement me")
}

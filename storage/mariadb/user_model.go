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
	"database/sql"
	"errors"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const DateTime string = "2006-01-02 15:04:05"

// UserModel is a storage.UserModel implementation.
type UserModel struct {
	DB *sqlx.DB
}

// Migrate implements storage.Model.Migrate.
func (u UserModel) Migrate() error {
	query := `
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    mail_addr VARCHAR(254) NOT NULL,
    name VARCHAR(50) NOT NULL,
    is_admin BOOLEAN NOT NULL,
    paypal_mail_addr VARCHAR(254) NOT NULL,
    score INTEGER NOT NULL,
    password_hash CHAR(60) NOT NULL,
    created DATETIME NOT NULL
)`

	_, err := u.DB.Exec(query)
	return err
}

// Drop implements storage.Model.Drop.
func (u UserModel) Drop() error {
	return drop(u.DB, "users")
}

// Create implements storage.UserModel.Create.
func (u UserModel) Create(user dl.User) error {
	query := `
INSERT INTO users (mail_addr, name, is_admin, paypal_mail_addr, score, password_hash, created)
VALUES (?, ?, ?, ?, ?, ?, ?)`

	created := user.Created.Format(DateTime)

	_, err := u.DB.Exec(query, user.MailAddr, user.Name, user.IsAdmin, user.PaypalMailAddr, user.Score, user.PasswordHash, created)
	if err != nil {
		return err
	}

	return nil
}

// Authenticate implements storage.UserModel.Authenticate.
func (u UserModel) Authenticate(mailAddr, password string) (int, error) {
	query := `SELECT id, password_hash FROM users WHERE mail_addr = ?`

	var id int
	var passwordHash []byte
	err := u.DB.QueryRowx(query, mailAddr).Scan(&id, &passwordHash)

	if err == sql.ErrNoRows {
		return 0, errors.New("invalid credentials")
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, errors.New("invalid credentials")
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

// FindByMailAddr implements storage.UserModel.FindByMailAddr.
func (u UserModel) FindByMailAddr(mailAddr string) (dl.User, error) {
	query := `SELECT * FROM users WHERE mail_addr = ?`

	var user dl.User
	err := u.DB.QueryRowx(query, mailAddr).StructScan(&user)

	return user, err
}

// Exists implements storage.UserModel.Exists.
func (u UserModel) Exists(mailAddr string) (bool, error) {
	query := `SELECT * FROM users where mail_addr = ?`

	var user dl.User
	err := u.DB.QueryRowx(query, mailAddr).StructScan(&user)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

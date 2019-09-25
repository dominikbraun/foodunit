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

// Package model provides basic domain models.
package model

import "time"

type User struct {
	ID             uint64    `db:"id"`
	MailAddr       string    `db:"mail_addr"`
	Name           string    `db:"name"`
	IsAdmin        bool      `db:"is_admin"`
	PaypalMailAddr string    `db:"paypal_mail_addr"`
	Score          int       `db:"score"`
	PasswordHash   []byte    `db:"password_hash"`
	Created        time.Time `db:"created"`
}

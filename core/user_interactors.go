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

// Package core provides business logic interactors and services.
package core

import (
	"github.com/dominikbraun/foodunit/core/dto"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// RegisterUser performs an user registration based on the provided UserRegistration data.
func RegisterUser(registration dto.UserRegistration, model storage.UserModel) error {
	// ToDo: server-side input validations.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), 12)
	if err != nil {
		return err
	}

	exists, err := model.Exists(registration.MailAddr)
	if err != nil {
		return err
	}
	if exists {
		return errors.Wrap(err, "duplicate E-Mail address")
	}

	user := dl.User{
		MailAddr:       registration.MailAddr,
		Name:           registration.Name,
		IsAdmin:        false,
		PaypalMailAddr: registration.MailAddr,
		Score:          10,
		PasswordHash:   hashedPassword,
		Created:        time.Now(),
	}

	err = model.Create(user)
	return err
}

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

// Package user provides services and types for User-related data.
package user

import (
	"github.com/dominikbraun/foodunit/model"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrUserExists      = errors.New("the given mail address already exists")
	ErrPasswordInvalid = errors.New("the given password is invalid")
	ErrUserNotStored   = errors.New("the user could not be registered")
)

type Service struct {
	users storage.User
}

func NewService(u storage.User) *Service {
	service := Service{
		users: u,
	}
	return &service
}

func (s *Service) Register(r *Registration) (bool, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, ErrPasswordInvalid
	}

	mailExists, err := s.users.MailExists(r.MailAddr)

	if err != nil {
		return false, err
	} else if mailExists {
		return false, ErrUserExists
	}

	user := model.User{
		MailAddr:       r.MailAddr,
		Name:           r.Name,
		IsAdmin:        false,
		PaypalMailAddr: r.PaypalMailAddr,
		Score:          0,
		PasswordHash:   hashedPassword,
		Created:        time.Now(),
	}

	err = s.users.Store(&user)
	if err != nil {
		return false, ErrUserNotStored
	}

	return true, nil
}

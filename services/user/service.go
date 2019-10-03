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
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/dominikbraun/foodunit/config"
	"github.com/dominikbraun/foodunit/model"
	"github.com/dominikbraun/foodunit/services/mail"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrUserExists              = errors.New("the given mail address already exists")
	ErrPasswordInvalid         = errors.New("the given password is invalid")
	ErrPasswordIncorrect       = errors.New("the given password is not correct")
	ErrUserNotStored           = errors.New("the user could not be registered")
	ErrUserNotFound            = errors.New("the user could not be found")
	ErrConfirmationMailNotSent = errors.New("confirmation mail could not be sent")
)

type Service struct {
	appConfig   config.Reader
	users       storage.User
	mailService *mail.Service
}

func NewService(r config.Reader, u storage.User, m *mail.Service) *Service {
	service := Service{
		appConfig:   r,
		users:       u,
		mailService: m,
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

	userEntity := model.User{
		MailAddr:       r.MailAddr,
		Name:           r.Name,
		IsAdmin:        false,
		PaypalMailAddr: r.PaypalMailAddr,
		Score:          0,
		PasswordHash:   hashedPassword,
		Created:        time.Now(),
	}

	userID, err := s.users.Store(&userEntity)
	if err != nil {
		return false, ErrUserNotStored
	}

	token := s.generateToken(userEntity.MailAddr)
	if err = s.users.StoreConfirmationToken(userID, token); err != nil {
		return false, err
	}

	if err = s.sendConfirmationMail(userEntity.MailAddr); err != nil {
		return false, ErrConfirmationMailNotSent
	}
	return true, nil
}

func (s *Service) generateToken(mailAddr string) string {
	data := []byte(mailAddr)
	hash := md5.Sum(data)

	return fmt.Sprintf("%x", hash)
}

func (s *Service) sendConfirmationMail(mailAddr string) error {
	settings := mail.Settings{
		From:    s.appConfig.GetString("confirmation_mail_from"),
		Subject: s.appConfig.GetString("confirmation_mail_subject"),
		Body:    s.appConfig.GetString("confirmation_mail_body"),
	}

	err := s.mailService.Send(&settings)
	return err
}

func (s *Service) Authenticate(l *Login) (uint64, error) {
	userEntity, err := s.users.FindByMailAddr(l.MailAddr)

	if err == sql.ErrNoRows {
		return 0, ErrUserNotFound
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(userEntity.PasswordHash, []byte(l.Password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, ErrPasswordIncorrect
	} else if err != nil {
		return 0, err
	}

	return userEntity.ID, nil
}

func (s *Service) ConfirmMailAddr(token string) error {
	return nil
}

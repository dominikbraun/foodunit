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

const (
	confirmationMailFrom    string = "confirmation_mail_from"
	confirmationMailSubject string = "confirmation_mail_subject"
	confirmationMailBody    string = "confirmation_mail_body"
)

var (
	ErrUserExists              = errors.New("the given mail address already exists")
	ErrPasswordInvalid         = errors.New("the given password is invalid")
	ErrPasswordIncorrect       = errors.New("the given password is not correct")
	ErrUserNotStored           = errors.New("the user could not be registered")
	ErrUserNotFound            = errors.New("the user could not be found")
	ErrConfirmationMailNotSent = errors.New("confirmation mail could not be sent")
	ErrTokenInvalid            = errors.New("the confirmation token is invalid")
)

// Service executes user-related business logic and use cases. It is also responsible
// for accessing the model storage under consideration of all business rules.
type Service struct {
	appConfig   config.Reader
	users       storage.User
	mailService *mail.Service
}

// NewService creates a new Service instance utilizing the given storage objects.
// The storage objects need to be ready to use for the service.
func NewService(r config.Reader, u storage.User, m *mail.Service) *Service {
	service := Service{
		appConfig:   r,
		users:       u,
		mailService: m,
	}
	return &service
}

// Register stores a new user, creates a confirmation token and sends a corresponding
// confirmation e-mail.
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

	if err = s.sendConfirmationMail(userEntity.Name, userEntity.MailAddr, token); err != nil {
		return false, ErrConfirmationMailNotSent
	}
	return true, nil
}

// generateToken generates a simple confirmation token which will be used for
// confirming a new user account.
func (s *Service) generateToken(mailAddr string) string {
	data := []byte(mailAddr)
	hash := md5.Sum(data)

	return fmt.Sprintf("%x", hash)
}

// sendConfirmationMail sends the actual confirmation e-mail. The mail subject and
// body will be read from the global application configuration.
func (s *Service) sendConfirmationMail(name, mailAddr, token string) error {
	from := s.appConfig.GetString(confirmationMailFrom)
	subj := s.appConfig.GetString(confirmationMailSubject)
	body := s.appConfig.GetString(confirmationMailBody)

	settings := mail.Settings{
		From:    from,
		To:      mailAddr,
		ToName:  name,
		Subject: subj,
		Body:    body,
		Variables: map[string]string{
			"token": token,
		},
	}

	err := s.mailService.Send(&settings)
	return err
}

// Authenticate will authenticate a user by comparing given login data with the
// stored password hash. Authenticate is not responsible for creating a session,
// this will be done by the session manager if the authentication was successful.
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

// ConfirmMailAddr will confirm the mail address (and therefore the user account)
// that is associated with the given token.
func (s *Service) ConfirmMailAddr(token string) error {
	err := s.users.ConfirmUser(token)

	if err == sql.ErrNoRows {
		return ErrTokenInvalid
	} else if err != nil {
		return err
	}

	return nil
}

// SetPaypalMailAddr sets the PayPal mail address of the user identified by id.
func (s *Service) SetPaypalMailAddr(id uint64, setter PaypalMailAddrSetter) error {
	err := s.users.SetPaypalMailAddr(id, setter.PaypalMailAddr)

	if err == sql.ErrNoRows {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}

	return nil
}

// Get returns all relevant meta data for an user identified by id.
func (s *Service) Get(id uint64) (PublicUser, error) {
	userEntity, err := s.users.Find(id)

	if err == sql.ErrNoRows {
		return PublicUser{}, ErrUserNotFound
	} else if err != nil {
		return PublicUser{}, err
	}

	offerView := PublicUser{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}

	return offerView, nil
}

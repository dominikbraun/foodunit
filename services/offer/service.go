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

// Package offer provides services and types for Offer-related data.
package offer

import (
	"database/sql"
	"github.com/dominikbraun/foodunit/config"
	"github.com/dominikbraun/foodunit/model"
	"github.com/dominikbraun/foodunit/services/mail"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/pkg/errors"
	"log"
	"time"
)

const (
	cancellationMailFrom    string = "cancellation_mail_from"
	cancellationMailSubject string = "cancellation_mail_subject"
	cancellationMailBody    string = "cancellation_mail_body"
)

var (
	ErrRestaurantNotFound = errors.New("the restaurant could not be found")
	ErrUserNotFound       = errors.New("the user could not be found")
	ErrOfferNotFound      = errors.New("the offer could not be found")
	ErrActionNotAllowed   = errors.New("the action is not allowed")
)

type Service struct {
	appConfig   config.Reader
	restaurants storage.Restaurant
	users       storage.User
	offers      storage.Offer
	orders      storage.Order
	positions   storage.Position
	mailService *mail.Service
}

func NewService(r config.Reader, res storage.Restaurant, u storage.User, o storage.Offer, odr storage.Order, p storage.Position, m *mail.Service) *Service {
	service := Service{
		appConfig:   r,
		restaurants: res,
		users:       u,
		offers:      o,
		orders:      odr,
		positions:   p,
		mailService: m,
	}
	return &service
}

func (s *Service) Create(c *Creation, userID uint64) error {
	userEntity, err := s.users.Find(userID)

	if err == sql.ErrNoRows {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}

	restaurant, err := s.restaurants.Find(c.Restaurant)

	if err == sql.ErrNoRows {
		return ErrRestaurantNotFound
	} else if err != nil {
		return err
	}

	offer := model.Offer{
		Owner:       userEntity,
		Restaurant:  restaurant,
		ValidFrom:   c.ValidFrom,
		ValidTo:     c.ValidTo,
		Responsible: userEntity,
		IsPlaced:    false,
		// ToDo: Set this value to NULL since ReadyAt isn't known at this point
		ReadyAt:       time.Now(),
		PaypalEnabled: c.PaypalEnabled,
	}

	err = s.offers.Store(&offer)
	return err
}

func (s *Service) Active() ([]ActiveOffer, error) {
	offerEntities, err := s.offers.FindValidFrom(time.Now())

	activeOffers := make([]ActiveOffer, 0)

	if err == sql.ErrNoRows {
		return activeOffers, nil
	} else if err != nil {
		return nil, err
	}

	for _, o := range offerEntities {
		activeOffer := ActiveOffer{
			ID:            o.ID,
			Owner:         User{ID: o.Owner.ID},
			Restaurant:    Restaurant{ID: o.Restaurant.ID},
			ValidFrom:     o.ValidFrom,
			ValidTo:       o.ValidTo,
			PaypalEnabled: o.PaypalEnabled,
		}

		activeOffers = append(activeOffers, activeOffer)
	}

	return activeOffers, nil
}

func (s *Service) Get(id uint64) (View, error) {
	offerEntity, err := s.offers.Find(id)

	if err == sql.ErrNoRows {
		return View{}, ErrOfferNotFound
	} else if err != nil {
		return View{}, err
	}

	offerView := View{
		ID:            offerEntity.ID,
		Owner:         User{Name: offerEntity.Owner.Name},
		ValidFrom:     offerEntity.ValidFrom,
		ValidTo:       offerEntity.ValidTo,
		Responsible:   User{Name: offerEntity.Responsible.Name},
		IsPlaced:      offerEntity.IsPlaced,
		IsCancelled:   offerEntity.IsCancelled,
		ReadyAt:       offerEntity.ReadyAt,
		PaypalEnabled: offerEntity.PaypalEnabled,
	}

	return offerView, nil
}

func (s *Service) Cancel(offerID, userID uint64) error {
	offer, err := s.offers.Find(offerID)

	if err == sql.ErrNoRows {
		return ErrOfferNotFound
	} else if err != nil {
		return err
	}

	if offer.Owner.ID != userID {
		return ErrActionNotAllowed
	}

	err = s.offers.Cancel(offerID)

	if err == sql.ErrNoRows {
		return ErrOfferNotFound
	} else if err != nil {
		return err
	}

	orderingUsers, err := s.offers.OrderingUsers(offerID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err != sql.ErrNoRows {
		owner, err := s.users.Find(offer.Owner.ID)

		if err != nil && err == sql.ErrNoRows {
			return ErrUserNotFound
		} else if err != nil {
			return err
		}

		restaurant, err := s.restaurants.Find(offer.Restaurant.ID)
		if err != nil && err == sql.ErrNoRows {
			return ErrRestaurantNotFound
		} else if err != nil {
			return err
		}

		for _, u := range orderingUsers {
			err = s.sendCancellationMail(u.Name, u.MailAddr, owner.Name, restaurant.Name)
			if err != nil {
				log.Printf("Error while cancelling offer #{%v}: {%s}\n", offerID, err.Error())
			}
		}
	}

	return nil
}

func (s *Service) sendCancellationMail(name, mailAddr, owner, restaurant string) error {
	from := s.appConfig.GetString(cancellationMailFrom)
	subj := s.appConfig.GetString(cancellationMailSubject)
	body := s.appConfig.GetString(cancellationMailBody)

	settings := mail.Settings{
		From:    from,
		To:      mailAddr,
		ToName:  name,
		Subject: subj,
		Body:    body,
		Variables: map[string]string{
			"owner":      owner,
			"restaurant": restaurant,
		},
	}

	err := s.mailService.Send(&settings)
	return err
}

func (s *Service) SetReadyAt(id uint64, setter ReadyAtSetter) error {
	err := s.offers.SetReadyAt(id, setter.ReadyAt)

	if err == sql.ErrNoRows {
		return ErrOfferNotFound
	} else if err != nil {
		return err
	}

	return nil
}

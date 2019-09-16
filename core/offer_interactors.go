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
	"time"
)

// CreateOffer creates a new offer based on the provided data
func CreateOffer(newOffer dto.NewOffer, offers storage.OfferModel, users storage.UserModel, restaurants storage.RestaurantModel) error {

	exists, err := restaurants.Exists(newOffer.Restaurant)
	if err != nil {
		return err
	}
	if !exists {
		return errors.Errorf("restaurant doesn't exist: %v", newOffer.Restaurant)
	}

	exists, err = users.Exists(newOffer.Responsible)
	if err != nil {
		return err
	}
	if !exists {
		return errors.Errorf("the responsible user doesn't exist: %v", newOffer.Responsible)
	}

	offer := dl.Offer{
		Owner:         dl.NewUser(0), // ToDo: get current user?
		Restaurant:    dl.NewRestaurant(newOffer.Restaurant),
		ValidFrom:     newOffer.ValidFrom,
		ValidTo:       newOffer.ValidTo,
		Responsible:   dl.NewUser(newOffer.Responsible),
		IsPlaced:      newOffer.IsPlaced,
		ReadyAt:       newOffer.ReadyAt,
		PaypalEnabled: newOffer.PaypalEnabled,
	}

	err = offers.Create(offer)
	return err
}

// GetActiveOffers delivers all offers which are active
func GetActiveOffers(offers storage.OfferModel) (dto.Offers, error) {
	offerIds, err := offers.GetActive(time.Now())

	if err != nil {
		return dto.Offers{}, err
	}

	return dto.Offers{Offers: offerIds}, nil
}

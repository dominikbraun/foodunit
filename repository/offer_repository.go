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

// Package repository provides discrete DL repository implementations.
package repository

import (
	"github.com/dominikbraun/foodunit/dl"
	"github.com/dominikbraun/foodunit/storage"
)

// UserRepository represents the default repository for managing Users.
type OfferRepository struct {
	storage storage.Repository
}

// Create implements dl.OfferRepository.Create.
func (o OfferRepository) Create(offer dl.Offer) error {
	err := o.storage.Create(offer)
	return err
}

// Find implements dl.OfferRepository.Find.
func (o OfferRepository) Find(id uint64) (dl.Offer, error) {
	var offer dl.Offer
	err := o.storage.Find(id, &offer)

	return offer, err
}

// Update implements dl.OfferRepository.Update.
func (o OfferRepository) Update(offer dl.Offer) error {
	err := o.storage.Update(offer)
	return err
}

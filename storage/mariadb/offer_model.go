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
	"github.com/jmoiron/sqlx"
	"time"
)

// OfferModel is a storage.OfferModel implementation.
type OfferModel struct {
	DB *sqlx.DB
}

// Migrate implements storage.Model.Migrate.
func (o OfferModel) Migrate() error {
	query := `
CREATE TABLE offers (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	owner_user_id BIGINT UNSIGNED NOT NULL,
	restaurant_id BIGINT UNSIGNED NOT NULL,
	valid_from DATETIME NOT NULL,
	valid_to DATETIME NOT NULL,
	responsible_user_id BIGINT UNSIGNED NOT NULL,
	is_placed BOOLEAN NOT NULL,
	ready_at DATETIME NOT NULL,
	paypal_enabled BOOLEAN NOT NULL
)`
	_, err := exec(o.DB, query)
	return err
}

// Drop implements storage.Model.Drop.
func (o OfferModel) Drop() error {
	return drop(o.DB, "offers")
}

// Create implements storage.OfferModel.Create.
func (o OfferModel) Create(offer dl.Offer) error {
	query := `
INSERT INTO offers (owner_user_id, restaurant_id, valid_from, valid_to, responsible_user_id, is_placed, ready_at, paypal_enabled)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	validFrom := offer.ValidFrom.Format(DateTime)
	validTo := offer.ValidTo.Format(DateTime)
	readyAt := offer.ReadyAt.Format(DateTime)

	_, err := exec(o.DB, query, offer.Owner.ID, offer.Restaurant.ID, validFrom, validTo, offer.Responsible.ID, offer.IsPlaced, readyAt, offer.PaypalEnabled)
	if err != nil {
		return err
	}

	return nil
}

// GetActive implements storage.OfferModel.GetActive.
func (o OfferModel) GetActive(now time.Time) ([]uint64, error) {
	query := `SELECT id FROM offers WHERE valid_from <= ? and ? <= valid_to`

	rows, err := o.DB.Queryx(query, now, now)
	if err != nil {
		return nil, err
	}

	ids := make([]uint64, 0)

	for rows.Next() {
		var offer dl.Offer

		err = rows.StructScan(&offer)
		if err != nil {
			return nil, err
		}

		ids = append(ids, offer.ID)
	}

	return ids, nil
}

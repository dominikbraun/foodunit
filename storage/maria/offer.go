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

// Package maria provides storage implementations for MariaDB.
package maria

import (
	"github.com/dominikbraun/foodunit/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type Offer struct {
	DB *sqlx.DB
}

func NewOffer(db *sqlx.DB) *Offer {
	offer := Offer{
		DB: db,
	}
	return &offer
}

func (o *Offer) Create() error {
	query := `
CREATE TABLE offers (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	owner_user_id BIGINT UNSIGNED NOT NULL,
	restaurant_id BIGINT UNSIGNED NOT NULL,
	valid_from DATETIME NOT NULL,
	valid_to DATETIME NOT NULL,
	responsible_user_id BIGINT UNSIGNED NOT NULL,
	is_placed BOOLEAN NOT NULL,
	is_cancelled BOOLEAN NOT NULL,
	ready_at DATETIME NOT NULL,
	paypal_enabled BOOLEAN NOT NULL
)`

	_, err := o.DB.Exec(query)
	return err
}

func (o *Offer) Drop() error {
	query := `DROP TABLE IF EXISTS offers`
	_, err := o.DB.Exec(query)

	return err
}

func (o *Offer) Store(offer *model.Offer) error {
	query := `
INSERT INTO offers (
	owner_user_id, restaurant_id, valid_from, valid_to,
	responsible_user_id, is_placed, is_cancelled, ready_at, paypal_enabled
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	validFrom := offer.ValidFrom.Format("2006-01-02 15:04:05")
	validTo := offer.ValidTo.Format("2006-01-02 15:04:05")

	// ToDo: Make ready_at column nullable and handle NULL in model
	readyAt := offer.ReadyAt.Format("2006-01-02 15:04:05")

	_, err := o.DB.Exec(
		query, offer.Owner.ID, offer.Restaurant.ID, validFrom,
		validTo, offer.Responsible.ID, offer.IsPlaced, offer.IsCancelled,
		readyAt, offer.PaypalEnabled,
	)

	return err
}

func (o *Offer) Find(id uint64) (model.Offer, error) {
	query := `
SELECT o.id, valid_from, valid_to, is_placed, is_cancelled, ready_at, paypal_enabled,
	u.id as "owner_user_id.id", u.name as "owner_user_id.name",
	r.id as "restaurant_id.id", r.name as "restaurant_id.name",
	u2.id as "responsible_user_id.id", u2.name as "responsible_user_id.name"
FROM offers o
INNER JOIN users u
ON u.id = o.owner_user_id
INNER JOIN restaurants r
ON r.id = o.restaurant_id
INNER JOIN users u2
ON u2.id = o.responsible_user_id
WHERE o.id = ?`

	var offer model.Offer
	err := o.DB.QueryRowx(query, id).StructScan(&offer)

	return offer, err
}

func (o *Offer) FindValidFrom(from time.Time) ([]model.Offer, error) {
	query := `
SELECT o.id, valid_from, valid_to, is_placed, is_cancelled, ready_at, paypal_enabled,
	u.id as "owner_user_id.id", u.name as "owner_user_id.name",
	r.id as "restaurant_id.id", r.name as "restaurant_id.name",
	u2.id as "responsible_user_id.id", u2.name as "responsible_user_id.name"
FROM offers o
INNER JOIN users u
ON u.id = o.owner_user_id
INNER JOIN restaurants r
ON r.id = o.restaurant_id
INNER JOIN users u2
ON u2.id = o.responsible_user_id
WHERE valid_from >= ?
AND is_placed = 0`

	validFrom := from.Format("2006-01-02 15:04:05")

	rows, err := o.DB.Queryx(query, validFrom)
	if err != nil {
		return nil, err
	}

	offers := make([]model.Offer, 0)

	for rows.Next() {
		var offer model.Offer

		if err := rows.StructScan(&offer); err != nil {
			return nil, err
		}
		offers = append(offers, offer)
	}

	return offers, nil
}

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

// ToDo: Don't forget Valid-To!
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
WHERE valid_from <= ?
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

func (o *Offer) Cancel(id uint64) error {
	query := `UPDATE offers SET is_cancelled = 1 WHERE id = ?`
	_, err := o.DB.Exec(query, id)

	return err
}

func (o *Offer) OwnerID(id uint64) (uint64, error) {
	query := `SELECT owner_user_id FROM offers WHERE id = ?`

	var ownerID int
	err := o.DB.QueryRowx(query, id).Scan(&ownerID)

	return uint64(ownerID), err
}

func (o *Offer) SetReadyAt(id uint64, readyAt time.Time) error {
	query := `UPDATE offers SET ready_at = ? WHERE id = ?`
	_, err := o.DB.Exec(query, readyAt, id)

	return err
}

func (o *Offer) OrderingUsers(id uint64) ([]model.User, error) {
	query := `
SELECT u.id as "id", u.mail_addr as "mail_addr"
FROM orders o
INNER JOIN users u
ON o.user_id = u.id
WHERE o.offer_id = ?`

	rows, err := o.DB.Queryx(query, id)
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)

	for rows.Next() {
		var user model.User

		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

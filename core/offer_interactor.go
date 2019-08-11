// Package core provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import (
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/pkg/errors"
)

// PositionMap holds multiple positions in an user's order. The key
// represents the DishID, while the value depicts the Note.
type PositionMap map[uint64]string

// OfferInteractor represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type OfferInteractor struct {
	offers    dl.OfferRepository
	orders    dl.OrderRepository
	positions dl.PositionRepository
}

// Creates a new OfferInteractor instance.
func NewOfferInteractor() (*OfferInteractor, error) {
	loader := load.RepositoryLoader

	if !loader.IsReady {
		return nil, errors.New("repository loader has to be initialized first")
	}

	return &OfferInteractor{
		offers:    loader.Offers(),
		orders:    loader.Orders(),
		positions: loader.Positions(),
	}, nil
}

// SaveOrder persists a new order that contains the given positions.
func (o *OfferInteractor) SaveOrder(userID, offerID uint64, pos PositionMap) error {
	order := dl.Order{
		UserID:  userID,
		OfferID: offerID,
	}

	orderID, err := o.orders.Create(&order)
	if err != nil {
		return err
	}

	for id, note := range pos {
		position := dl.Position{
			ID:      0,
			OrderID: orderID,
			DishID:  id,
			Note:    note,
		}
		// ToDo: The error value of Create has to be handled.
		_, _ = o.positions.Create(&position)
	}
	return nil
}

// GetOrders returns all saved orders for a given offer.
func (o *OfferInteractor) GetOrders(offerID uint64) ([]*dl.Order, error) {
	orders, err := o.orders.FindByOfferID(offerID)
	return orders, err
}

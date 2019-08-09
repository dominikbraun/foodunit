// Package `core` provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import (
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/dl"
	"github.com/pkg/errors"
)

// `OfferInteractor` represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type OfferInteractor struct {
	offers    dl.OfferRepository
	orders    dl.OrderRepository
	positions dl.PositionRepository
}

// Creates a new `OfferInteractor` instance.
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

// `GetOrders` returns all saved orders for a given offer.
func (o *OfferInteractor) GetOrders(offerID uint64) []*dl.Order {
	orders := o.orders.FindByOfferID(offerID)
	return orders
}

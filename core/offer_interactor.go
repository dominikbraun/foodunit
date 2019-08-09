// Package `core` provides types for interacting with the core business
// logic. It depicts the default use case of FoodUnit.
package core

import (
	"github.com/dominikbraun/foodunit/core/load"
	"github.com/dominikbraun/foodunit/dl"
)

// `OfferInteractor` represents an interface to be used by the adapters
// for triggering the core business logic and receiving the results.
type OfferInteractor struct {
	offers    dl.OfferRepository
	orders    dl.OrderRepository
	positions dl.PositionRepository
}

// Creates a new `OfferInteractor` instance.
func NewOfferInteractor() *OfferInteractor {
	return &OfferInteractor{
		offers:    load.OfferRepository(),
		orders:    load.OrderRepository(),
		positions: load.PositionRepository(),
	}
}

func (o *OfferInteractor) GetOrders(offerID uint64) []*dl.Order {
	orders := o.orders.FindByOfferID(offerID)
	return orders
}

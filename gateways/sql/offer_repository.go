// package `sql` provides repository implementations as SQL gateways.
package sql

import (
	"github.com/dominikbraun/foodunit/dl"
)

type (
	// `OfferRepository` holds methods of `dl.OfferRepository`.
	OfferRepository struct{}
	// `OrderRepository` holds methods of `dl.OrderRepository`.
	OrderRepository struct{}
	// `PositionRepository` holds methods of `dl.PositionRepository`.
	PositionRepository struct{}
)

// `Migrate` implements `dl.OfferRepository.`.
func (o OfferRepository) Migrate() error {
	panic("implement me")
}

// `Create` implements `dl.OfferRepository.Create`.
func (o OfferRepository) Create(offer *dl.Offer) (uint64, error) {
	panic("implement me")
}

// `Find` implements `dl.OfferRepository.Find`.
func (o OfferRepository) Find(id uint64) *dl.Offer {
	panic("implement me")
}

// `FindAllActive` implements `dl.OfferRepository.FindAllActive`.
func (o OfferRepository) FindAllActive() []*dl.Offer {
	panic("implement me")
}

// `Update` implements `dl.OfferRepository.Update`.
func (o OfferRepository) Update(offer *dl.Offer) error {
	panic("implement me")
}

// `Delete` implements `dl.OfferRepository.Delete`.
func (o OfferRepository) Delete(offer *dl.Offer) error {
	panic("implement me")
}

// `Migrate` implements `dl.OrderRepository.Migrate`.
func (o OrderRepository) Migrate() error {
	panic("implement me")
}

// `Create` implements `dl.OrderRepository.Create`.
func (o OrderRepository) Create(order *dl.Order) (uint64, error) {
	panic("implement me")
}

// `Find` implements `dl.OrderRepository.Find`.
func (o OrderRepository) Find(id uint64) (*dl.Order, error) {
	panic("implement me")
}

// `FindByOfferID` implements `dl.OrderRepository.FindByOfferID`.
func (o OrderRepository) FindByOfferID(offerID uint64) ([]*dl.Order, error) {
	panic("implement me")
}

// `Update` implements `dl.OrderRepository.Update`.
func (o OrderRepository) Update(order *dl.Order) error {
	panic("implement me")
}

// `Delete` implements `dl.OrderRepository.Delete`.
func (o OrderRepository) Delete(order *dl.Order) error {
	panic("implement me")
}

// `Migrate` implements `dl.PositionRepository.Migrate`.
func (p PositionRepository) Migrate() error {
	panic("implement me")
}

// `Create` implements `dl.PositionRepository.Create`.
func (p PositionRepository) Create(position *dl.Position) (uint64, error) {
	panic("implement me")
}

// `Find` implements `dl.PositionRepository.Find`.
func (p PositionRepository) Find(id uint64) (*dl.Position, error) {
	panic("implement me")
}

// `FindByOrderID` implements `dl.PositionRepository.FindByOrderID`.
func (p PositionRepository) FindByOrderID(orderID uint64) ([]*dl.Position, error) {
	panic("implement me")
}

// `Update` implements `dl.PositionRepository.Update`.
func (p PositionRepository) Update(position *dl.Position) error {
	panic("implement me")
}

// `Delete` implements `dl.PositionRepository.Delete`.
func (p PositionRepository) Delete(position *dl.Position) error {
	panic("implement me")
}

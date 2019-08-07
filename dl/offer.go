// Package dl provides Domain Language entities and rules.
package dl

import "time"

// `Offer` represents an `User`'s offer to order food for their team.
type Offer struct {
	ID         uint64
	UserID     uint64
	SupplierID uint64
	From       time.Time
	To         time.Time
	IsPlaced   bool
	PickupInfo string
}

// `OfferRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type OfferRepository interface {
	Create(o *Offer) error
	Find(id uint64) *Offer
	Update(o *Offer) error
	Delete(o *Offer) error
}

// `Order` represents an `User`'s order that was placed as part of
// an `Offer`. The list of ordered food is depicted in `Positions`.
type Order struct {
	ID      uint64
	UserID  uint64
	OfferID uint64
}

// `OrderRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type OrderRepository interface {
	Create(o *Order) error
	Find(id uint64) *Order
	Update(o *Order) error
	Delete(o *Order) error
}

// `Position` represents one of multiple items contained in an `Order`.
type Position struct {
	ID      uint64
	OrderID uint64
	DishID  uint64
	Note    string
}

// `PositionRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type PositionRepository interface {
	Create(p *Position) error
	Find(id uint64) *Position
	Update(p *Position) error
	Delete(p *Position) error
}

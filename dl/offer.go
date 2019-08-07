// Package dl provides Domain Language entities and rules.
package dl

import "time"

// `Offer` represents an `User`'s offer to order food for their team.
type Offer struct {
	ID         uint
	UserID     uint
	SupplierID uint
	From       time.Time
	To         time.Time
	IsPlaced   bool
	PickupInfo string
}

// `OfferRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type OfferRepository interface {
	Create(u *Offer) error
	Find(id uint) *Offer
	Update(u *Offer) error
	Delete(u *Offer) error
}

// `Order` represents an `User`'s order that was placed as part of
// an `Offer`. The list of ordered food is depicted in `Positions`.
type Order struct {
	ID      uint
	UserID  uint
	OfferID uint
}

// `OrderRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type OrderRepository interface {
	Create(u *Order) error
	Find(id uint) *Order
	Update(u *Order) error
	Delete(u *Order) error
}

// `Position` represents one of multiple items contained in an `Order`.
type Position struct {
	ID      uint
	OrderID uint
	DishID  uint
	Note    string
}

// `PositionRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type PositionRepository interface {
	Create(u *Position) error
	Find(id uint) *Position
	Update(u *Position) error
	Delete(u *Position) error
}

// Package dl provides Domain Language entities and rules.
package dl

import "time"

// Offer represents an User's offer to order food for their team.
type Offer struct {
	ID         uint64    `db:"id"`
	UserID     uint64    `db:"user_id"`
	SupplierID uint64    `db:"supplier_id"`
	ValidFrom  time.Time `db:"valid_from"`
	ValidTo    time.Time `db:"valid_to"`
	IsPlaced   bool      `db:"is_placed"`
	PickupInfo string    `db:"pickup_info"`
}

// OfferRepository provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type OfferRepository interface {
	Migrate() error
	Create(o *Offer) (uint64, error)
	Find(id uint64) *Offer
	FindAllActive() []*Offer
	Update(o *Offer) error
	Delete(o *Offer) error
}

// Order represents an User's order that was placed as part of
// an Offer. The list of ordered food is depicted in Positions.
type Order struct {
	ID      uint64 `db:"id"`
	UserID  uint64 `db:"user_id"`
	OfferID uint64 `db:"offer_id"`
}

// OrderRepository provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type OrderRepository interface {
	Migrate() error
	Create(o *Order) (uint64, error)
	Find(id uint64) (*Order, error)
	FindByOfferID(offerID uint64) ([]*Order, error)
	Update(o *Order) error
	Delete(o *Order) error
}

// Position represents one of multiple items contained in an Order.
type Position struct {
	ID      uint64 `db:"id"`
	OrderID uint64 `db:"order_id"`
	DishID  uint64 `db:"dish_id"`
	Note    string `db:"note"`
}

// PositionRepository provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type PositionRepository interface {
	Migrate() error
	Create(p *Position) (uint64, error)
	Find(id uint64) (*Position, error)
	FindByOrderID(orderID uint64) ([]*Position, error)
	Update(p *Position) error
	Delete(p *Position) error
}

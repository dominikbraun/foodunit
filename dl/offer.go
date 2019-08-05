package dl

import "time"

// DL description for `Offer`, `Order`, `Position` and their ORM repositories.

// `Offer` represents an `User`'s offer to order food for their team.
type Offer struct {
	ID         uint      `gorm:"primary_key"`
	CreatedBy  User      `gorm:"foreignkey:ID"`
	Supplier   Supplier  `gorm:"foreignkey:ID"`
	From       time.Time `gorm:""`
	To         time.Time `gorm:""`
	IsPlaced   bool      `gorm:""`
	PickupInfo string    `gorm:"type:varchar(100)"`
	Orders     []Order   `gorm:""`
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
	ID        uint       `gorm:"primary_key"`
	User      User       `gorm:"foreignkey:ID"`
	Positions []Position `gorm:""`
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
	ID      uint   `gorm:"primary_key"`
	OrderID uint   `gorm:""`
	DishID  uint   `gorm:""`
	Note    string `gorm:"type:varchar(50)"`
}

// `PositionRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type PositionRepository interface {
	Create(u *Position) error
	Find(id uint) *Position
	Update(u *Position) error
	Delete(u *Position) error
}

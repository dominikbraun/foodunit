// package `load` provides concrete interface implementations prescribed
// by the Domain Language.
package load

import "github.com/dominikbraun/foodunit/dl"

// `RepositoryLoader` provides functions for obtaining concrete repository
// implementations. This component has to be initialized by an interface
// adapter calling an use case.
type RepositoryLoader struct {
	IsReady    bool
	Users      func() dl.UserRepository
	Suppliers  func() dl.SupplierRepository
	Categories func() dl.CategoryRepository
	Dishes     func() dl.DishRepository
	Offers     func() dl.OfferRepository
	Orders     func() dl.OrderRepository
	Positions  func() dl.PositionRepository
}

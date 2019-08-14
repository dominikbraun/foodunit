// Package load provides concrete interface implementations prescribed
// by the Domain Language.
package load

import "github.com/dominikbraun/foodunit/dl"

// RepositoryLoader is the global instance which will be accessed by the
// interactors after it has been initialized by the corresponding adapter.
var RepositoryLoader *repositoryLoader

// repositoryLoader holds functions providing repository implementations.
type repositoryLoader struct {
	IsReady    bool
	Users      dl.UserRepository
	Suppliers  dl.SupplierRepository
	Categories dl.CategoryRepository
	Dishes     dl.DishRepository
	Offers     dl.OfferRepository
	Orders     dl.OrderRepository
	Positions  dl.PositionRepository
}

// init initializes the global repositoryLoader instance so that it
// can be used by an interface adapter out of the box.
func init() {
	RepositoryLoader = &repositoryLoader{}
	RepositoryLoader.IsReady = false
}

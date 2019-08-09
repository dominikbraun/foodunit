// package `load` provides concrete interface implementations prescribed
// by the Domain Language.
package load

import "github.com/dominikbraun/foodunit/dl"

// `UserRepository` loads a specific DL repository implementation.
func UserRepository() dl.UserRepository {
	return nil
}

// `SupplierRepository` loads a specific DL repository implementation.
func SupplierRepository() dl.SupplierRepository {
	return nil
}

// `CategoryRepository` loads a specific DL repository implementation.
func CategoryRepository() dl.CategoryRepository {
	return nil
}

// `DishRepository` loads a specific DL repository implementation.
func DishRepository() dl.DishRepository {
	return nil
}

// `OfferRepository` loads a specific DL repository implementation.
func OfferRepository() dl.OfferRepository {
	return nil
}

// `OrderRepository` loads a specific DL repository implementation.
func OrderRepository() dl.OrderRepository {
	return nil
}

// `PositionRepository` loads a specific DL repository implementation.
func PositionRepository() dl.PositionRepository {
	return nil
}

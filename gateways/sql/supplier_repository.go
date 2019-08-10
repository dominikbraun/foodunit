// package `sql` provides repository implementations as SQL gateways.
package sql

import (
	"github.com/dominikbraun/foodunit/dl"
)

type (
	// `SupplierRepository` holds methods of `dl.SupplierRepository`.
	SupplierRepository struct{}
	// `CategoryRepository` holds methods of `dl.CategoryRepository`.
	CategoryRepository struct{}
	// `DishRepository` holds methods of `dl.DishRepository`.
	DishRepository struct{}
)

// `Migrate` implements `dl.SupplierRepository.Migrate`.
func (s SupplierRepository) Migrate() error {
	panic("implement me")
}

// `Create` implements `dl.SupplierRepository.Create`.
func (s SupplierRepository) Create(supplier *dl.Supplier) (uint64, error) {
	panic("implement me")
}

// `Find` implements `dl.SupplierRepository.Find`.
func (s SupplierRepository) Find(id uint64) (*dl.Supplier, error) {
	panic("implement me")
}

// `Update` implements `dl.SupplierRepository.Update`.
func (s SupplierRepository) Update(supplier *dl.Supplier) error {
	panic("implement me")
}

// `Delete` implements `dl.SupplierRepository.Delete`.
func (s SupplierRepository) Delete(supplier *dl.Supplier) error {
	panic("implement me")
}

// `Migrate` implements `dl.CategoryRepository.Migrate`.
func (c CategoryRepository) Migrate() error {
	panic("implement me")
}

// `Create` implements `dl.CategoryRepository.Create`.
func (c CategoryRepository) Create(category *dl.Category) (uint64, error) {
	panic("implement me")
}

// `Find` implements `dl.CategoryRepository.Find`.
func (c CategoryRepository) Find(id uint64) (*dl.Category, error) {
	panic("implement me")
}

// `FindBySupplierID` implements `dl.CategoryRepository.FindBySupplierID`.
func (c CategoryRepository) FindBySupplierID(supplierID uint64) ([]*dl.Category, error) {
	panic("implement me")
}

// `Update` implements `dl.CategoryRepository.Update`.
func (c CategoryRepository) Update(category *dl.Category) error {
	panic("implement me")
}

// `Delete` implements `dl.CategoryRepository.Delete`.
func (c CategoryRepository) Delete(category *dl.Category) error {
	panic("implement me")
}

// `Migrate` implements `dl.DishRepository.Migrate`.
func (d DishRepository) Migrate() error {
	panic("implement me")
}

// `Create` implements `dl.DishRepository.Create`.
func (d DishRepository) Create(dish *dl.Dish) (uint64, error) {
	panic("implement me")
}

// `Find` implements `dl.DishRepository.Find`.
func (d DishRepository) Find(id uint64) (*dl.Dish, error) {
	panic("implement me")
}

// `FindByCategoryID` implements `dl.DishRepository.FindByCategoryID`.
func (d DishRepository) FindByCategoryID(categoryID uint64) ([]*dl.Dish, error) {
	panic("implement me")
}

// `Update` implements `dl.DishRepository.Update`.
func (d DishRepository) Update(dish *dl.Dish) error {
	panic("implement me")
}

// `Delete` implements `dl.DishRepository.Delete`.
func (d DishRepository) Delete(dish *dl.Dish) error {
	panic("implement me")
}

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
	schema := `
CREATE TABLE suppliers (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50),
	street VARCHAR(50),
	postal_code VARCHAR(10),
	city VARCHAR(25),
	open_mon VARCHAR(25),
	open_tue VARCHAR(25),
	open_wed VARCHAR(25),
	open_thu VARCHAR(25),
	open_fri VARCHAR(25),
	open_sat VARCHAR(25),
	open_sun VARCHAR(25),
	website VARCHAR(50)
)`

	db, err := GetDB()
	if err != nil {
		return err
	}

	if _, err = db.Exec(schema); err != nil {
		return err
	}
	return nil
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
	schema := `
CREATE TABLE categories (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	supplier_id BIGINT UNSIGNED,
	name VARCHAR(25),
	img_path VARCHAR(255)
)`

	db, err := GetDB()
	if err != nil {
		return err
	}

	if _, err = db.Exec(schema); err != nil {
		return err
	}
	return nil
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
	schema := `
CREATE TABLE dishes (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	category_id BIGINT UNSIGNED,
	name VARCHAR(25),
	description VARCHAR(50),
	price INT(8) UNSIGNED
)`

	db, err := GetDB()
	if err != nil {
		return err
	}

	if _, err = db.Exec(schema); err != nil {
		return err
	}
	return nil
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

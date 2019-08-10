// Package `dl` provides Domain Language entities and rules.
package dl

// `Supplier` represents a food supplier like a restaurant or delivery service.
type Supplier struct {
	ID         uint64 `db:"id"`
	Name       string `db:"name"`
	Street     string `db:"street"`
	PostalCode string `db:"postal_code"`
	City       string `db:"city"`
	OpenMon    string `db:"open_mon"`
	OpenTue    string `db:"open_tue"`
	OpenWed    string `db:"open_wed"`
	OpenThu    string `db:"open_thu"`
	OpenFri    string `db:"open_fri"`
	OpenSat    string `db:"open_sat"`
	OpenSun    string `db:"open_sun"`
	Website    string `db:"website"`
}

// `SupplierRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type SupplierRepository interface {
	Migrate() error
	Create(s *Supplier) (uint64, error)
	Find(id uint64) (*Supplier, error)
	Update(s *Supplier) error
	Delete(s *Supplier) error
}

// `Category` represents the category or menu's section a dish belongs to.
type Category struct {
	ID         uint64 `db:"id"`
	SupplierID uint64 `db:"supplier_id"`
	Name       string `db:"name"`
	ImgPath    string `db:"img_path"`
}

// `CategoryRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type CategoryRepository interface {
	Migrate() error
	Create(c *Category) (uint64, error)
	Find(id uint64) (*Category, error)
	FindBySupplierID(supplierID uint64) ([]*Category, error)
	Update(c *Category) error
	Delete(c *Category) error
}

// `Dish` represents a meal or food in general that is offered by a `Supplier`,
// where it is listed as a menu item.
type Dish struct {
	ID          uint64 `db:"id"`
	CategoryID  uint64 `db:"category_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       uint8  `db:"price"`
}

// `DishRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type DishRepository interface {
	Migrate() error
	Create(d *Dish) (uint64, error)
	Find(id uint64) (*Dish, error)
	FindByCategoryID(categoryID uint64) ([]*Dish, error)
	Update(d *Dish) error
	Delete(d *Dish) error
}

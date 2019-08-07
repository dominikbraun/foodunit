// Package dl provides Domain Language entities and rules.
package dl

// `Supplier` represents a food supplier like a restaurant or delivery service.
type Supplier struct {
	ID         uint64
	Name       string
	Street     string
	PostalCode string
	City       string
	OpenMon    string
	OpenTue    string
	OpenWed    string
	OpenThu    string
	OpenFri    string
	OpenSat    string
	OpenSun    string
	Website    string
}

// `SupplierRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type SupplierRepository interface {
	Create(s *Supplier) error
	Find(id uint64) *Supplier
	Update(s *Supplier) error
	Delete(s *Supplier) error
}

// `Category` represents the category or menu's section a dish belongs to.
type Category struct {
	ID         uint64
	SupplierID uint64
	Name       string
	ImgPath    string
}

// `CategoryRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type CategoryRepository interface {
	Create(c *Category) error
	Find(id uint64) *Category
	Update(c *Category) error
	Delete(c *Category) error
}

// `Dish` represents a meal or food in general that is offered by a `Supplier`,
// where it is listed as a menu item.
type Dish struct {
	ID          uint64
	CategoryID  uint64
	Name        string
	Description string
	Price       uint8
}

// `DishRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type DishRepository interface {
	Create(d *Dish) error
	Find(id uint64) *Dish
	Update(d *Dish) error
	Delete(d *Dish) error
}

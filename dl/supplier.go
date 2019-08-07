// Package dl provides Domain Language entities and rules.
package dl

// `Supplier` represents a food supplier like a restaurant or delivery service.
type Supplier struct {
	ID         uint
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
	Menu       []Category
}

// `SupplierRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type SupplierRepository interface {
	Create(u *Supplier) error
	Find(id uint) *Supplier
	Update(u *Supplier) error
	Delete(u *Supplier) error
}

// `Category` represents the category or menu's section a dish belongs to.
type Category struct {
	ID         uint
	SupplierID uint
	Name       string
	ImgPath    string
	Dishes     []Dish
}

// `CategoryRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type CategoryRepository interface {
	Create(u *Category) error
	Find(id uint) *Category
	Update(u *Category) error
	Delete(u *Category) error
}

// `Dish` represents a meal or food in general that is offered by a `Supplier`,
// where it is listed as a menu item.
type Dish struct {
	ID          uint
	CategoryID  uint
	Name        string
	Description string
	Price       uint
}

// `DishRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type DishRepository interface {
	Create(u *Dish) error
	Find(id uint) *Dish
	Update(u *Dish) error
	Delete(u *Dish) error
}

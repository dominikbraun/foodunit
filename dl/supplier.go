// Package dl provides Domain Language entities and rules.
package dl

// `Supplier` represents a food supplier like a restaurant or delivery service.
type Supplier struct {
	ID         uint       `gorm:"primary_key"`
	Name       string     `gorm:"type:varchar(100)"`
	Street     string     `gorm:"type:varchar(100)"`
	PostalCode string     `gorm:"type:varchar(50)"`
	City       string     `gorm:"type:varchar(100)"`
	OpenMon    string     `gorm:"type:varchar(50)"`
	OpenTue    string     `gorm:"type:varchar(50)"`
	OpenWed    string     `gorm:"type:varchar(50)"`
	OpenThu    string     `gorm:"type:varchar(50)"`
	OpenFri    string     `gorm:"type:varchar(50)"`
	OpenSat    string     `gorm:"type:varchar(50)"`
	OpenSun    string     `gorm:"type:varchar(50)"`
	Website    string     `gorm:"type:varchar(253)"`
	Menu       []Category `gorm:"foreignkey:SupplierID"`
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
	ID         uint   `gorm:"primary_key"`
	SupplierID uint   `gorm:""`
	Name       string `gorm:"type:varchar(50)"`
	ImgPath    string `gorm:"type:varchar(255)"`
	Dishes     []Dish `gorm:"foreignkey:CategoryID"`
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
	ID          uint   `gorm:"primary_key"`
	CategoryID  uint   `gorm:""`
	Name        string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:varchar(100)"`
	Price       uint   `gorm:"type:varchar(50)"`
}

// `DishRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type DishRepository interface {
	Create(u *Dish) error
	Find(id uint) *Dish
	Update(u *Dish) error
	Delete(u *Dish) error
}

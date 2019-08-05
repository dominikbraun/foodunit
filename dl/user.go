package dl

// DL description for `User` and its ORM repository.

// `User` represents a person which can login, create offers and order food.
type User struct {
	ID   uint   `gorm:"primary_key"`
	Mail string `gorm:"type:varchar(254);unique"`
	Name string `gorm:"type:varchar(100)"`
}

// `UserRepository` provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type UserRepository interface {
	Create(u *User) error
	Find(id uint) *User
	FindByMail(mail string) *User
	Update(u *User) error
	Delete(u *User) error
}

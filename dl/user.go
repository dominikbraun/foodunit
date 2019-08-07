// Package dl provides Domain Language entities and rules.
package dl

// `User` represents a person which can login, create offers and order food.
type User struct {
	ID   uint64 `db:"id"`
	Mail string `db:"mail"`
	Name string `db:"name"`
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

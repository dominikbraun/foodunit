// Package dl provides Domain Language entities and rules.
package dl

// User represents a person which can login, create offers and order food.
type User struct {
	ID   uint64 `db:"id"`
	Mail string `db:"mail"`
	Name string `db:"name"`
}

// UserRepository provides methods for typical CRUD operations. Its
// implementations are stored in the /gateways package.
type UserRepository interface {
	Migrate() error
	Create(u *User) (uint64, error)
	Find(id uint) (*User, error)
	FindByMail(mail string) (*User, error)
	Update(id uint64, u *User) error
	Delete(u *User) error
}

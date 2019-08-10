// package `sql` provides repository implementations as SQL gateways.
package sql

import (
	"github.com/dominikbraun/foodunit/dl"
)

type (
	// `Repository` holds methods of `dl.Repository`.
	UserRepository struct{}
)

// `Migrate` implements `dl.UserRepository.Migrate`.
func (u UserRepository) Migrate() error {
	panic("implement me")
}

// `Create` implements `dl.UserRepository.Create`.
func (u UserRepository) Create(user *dl.User) (uint64, error) {
	panic("implement me")
}

// `Find` implements `dl.UserRepository.Find`.
func (u UserRepository) Find(id uint) (*dl.User, error) {
	panic("implement me")
}

// `FindByMail` implements `dl.UserRepository.FindByMail`.
func (u UserRepository) FindByMail(mail string) (*dl.User, error) {
	panic("implement me")
}

// `Update` implements `dl.UserRepository.Update`.
func (u UserRepository) Update(user *dl.User) error {
	panic("implement me")
}

// `Delete` implements `dl.UserRepository.Delete`.
func (u UserRepository) Delete(user *dl.User) error {
	panic("implement me")
}

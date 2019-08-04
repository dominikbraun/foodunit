package user

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

// User represents a person that can create offers and
// order food. Each user has to register theirselves first.
type User struct {
	gorm.Model
	ID    int    `gorm:"primary_key"`
	Email string `gorm:"varchar(254);unique_index"`
	Name  string `gorm:"varchar(50)"`
}

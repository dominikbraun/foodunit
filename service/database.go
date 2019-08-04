package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	defaultCharset string = "utf8"
)

var cachedInstance *gorm.DB = nil

func NewDatabase(dialect, user, pass, db string) (*gorm.DB, error) {
	if cachedInstance == nil {
		conn := fmt.Sprintf("%s:%s@/%s?charset=%s", user, pass, db, defaultCharset)

		gdb, err := gorm.Open(dialect, conn)
		if err != nil {
			return nil, err
		}

		cachedInstance = gdb
	}

	return cachedInstance, nil
}

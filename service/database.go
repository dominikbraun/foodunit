package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	charset string = "utf8"
)

var cachedDB *gorm.DB = nil

func DB() (*gorm.DB, error) {
	if cachedDB == nil {

		c, err := Conf()
		if err != nil {
			return nil, err
		}

		user := c.Get("dbuser")
		pass := c.Get("dbpass")
		dbname := c.Get("dbname")

		conn := fmt.Sprintf("%s:%s@/%s?charset=%s", user, pass, dbname, charset)

		db, err := gorm.Open(c.Get("dbsys").(string), conn)
		if err != nil {
			return nil, err
		}
		cachedDB = db
	}
	return cachedDB, nil
}

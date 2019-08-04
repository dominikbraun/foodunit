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

		conn := fmt.Sprintf("%s:%s@/%s?charset=%s", c.Get("dbuser"), c.Get("dbpass"), c.Get("dbname"), charset)

		db, err := gorm.Open(c.Get("dbsys").(string), conn)
		if err != nil {
			return nil, err
		}
		cachedDB = db
	}
	return cachedDB, nil
}

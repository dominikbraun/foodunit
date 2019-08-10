package sql

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var instance *sqlx.DB = nil

func Connect(driver, dbname, user, pass string) error {
	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)

	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		return err
	}

	instance = db
	return nil
}

func GetDB() (*sqlx.DB, error) {
	if instance == nil {
		return nil, errors.New("invalid database instance - call `Connect` first")
	}
	return instance, nil
}

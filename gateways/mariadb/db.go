// package mariadb provides repository implementations as SQL gateways.
package mariadb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// instance represents the database connection to be used by all repositories.
// It has to be initialized with Connect by the gateway or framework which
// utilizes this instance first.
// This object should be exclusively accessed via GetDB().
var instance *sqlx.DB = nil

// Connect establishes a database connection using the given parameters. It
// will initialize the global database instance used by all repositories.
func Connect(driver, dbname, user, pass string) error {
	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)

	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		return err
	}

	instance = db
	return nil
}

// GetDB returns the global database instance. It is recommended to call
// this function instead of accessing the instance directly since it checks
// if the database connection has been established yet.
func GetDB() (*sqlx.DB, error) {
	if instance == nil {
		return nil, errors.New("invalid database instance - call `Connect` first")
	}
	return instance, nil
}

// lastInsertID returns the ID of a new inserted row. This function reduces
// repetitive error handling and casts to uint64 in repository implementations.
func lastInsertID(result sql.Result) (uint64, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), err
}

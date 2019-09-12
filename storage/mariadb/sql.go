package mariadb

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// drop attempts to delete a database table.
func drop(db *sqlx.DB, table string) error {
	query := `DROP TABLE IF EXISTS ` + table
	// ToDo: check for affected rows? (return error if 0?)
	_, err := exec(db, query)
	return err
}

// exe wraps DB.Exec and adds more error information in case of an error.
func exec(db *sqlx.DB, query string, args ...interface{}) (sql.Result, error) {
	res, err := db.Exec(query, args...)
	return res, errors.Wrap(err, query)
}

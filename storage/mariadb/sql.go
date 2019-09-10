package mariadb

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// drops a table
func drop(db *sqlx.DB, table string) error {
	query := `drop table if exists ` + table
	// ToDo: check for affected rows? (return error if 0?)
	_, err := exec(db, query)
	return err
}

// wraps DB.Exec and adds more error information in case of an error
func exec(db *sqlx.DB, query string, args ...interface{}) (sql.Result, error) {
	res, err := db.Exec(query)
	return res, errors.Wrap(err, query)
}

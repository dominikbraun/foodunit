package mariadb

import "github.com/jmoiron/sqlx"

func drop(db *sqlx.DB, table string) error {
	query := `drop table if exists ` + table
	// ToDo: check for affected rows? (return error if 0?)
	_, err := db.Exec(query)
	return err
}

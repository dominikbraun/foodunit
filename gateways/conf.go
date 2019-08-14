// Package gateways provides all DL repository implementations.
package gateways

// Conf represents a database configuration. It serves as an interchangeable
// data type which may be passed around from the user interface to the DB.
type Conf struct {
	Driver   string `json:"db_driver"`
	Host     string `json:"db_host"`
	DBName   string `json:"db_dbname"`
	User     string `json:"db_user"`
	Password string `json:"db_pass"`
}

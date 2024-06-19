package src

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func StartDB(dataSource string) (*sql.DB, error) {
	return sql.Open("sqlite3", dataSource)
}

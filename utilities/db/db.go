package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Start(dataSource string) (*sql.DB, error) {
	return sql.Open("sqlite3", dataSource)
}

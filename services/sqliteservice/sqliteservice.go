package sqliteservice

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLite() (*sql.DB, error) {
	return sql.Open("sqlite3", "server.db")
}

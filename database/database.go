package database

import (
	"database/sql"

	e "github.com/dehwyy/dehwyy-cli/error-handler"
)

func InitDB() *sql.DB {
	db := e.WithFatal(sql.Open("sqlite3", "./db.sqlite"))("Cannot open database")

	return db
}

func CloseDB(db *sql.DB) {
	e.WithFatalString(db.Close(), "Cannot close database")
}

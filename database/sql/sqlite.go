package sql_database

import (
	"database/sql"
	"log"
)

type Sqlite struct{}

func (s *Sqlite) New() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("Cannot open database: %v", err)
	}

	return db
}

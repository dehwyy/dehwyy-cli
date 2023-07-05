package database

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, e := sql.Open("sqlite3", "./db.sqlite")

	if e != nil {
			log.Fatalf("Cannot open database: %v", e)
	}

	return db
}

func CloseDB(db *sql.DB) {
	e := db.Close()
	if e != nil {
		log.Fatalf("Cannot close database: %v", e)
	}
}

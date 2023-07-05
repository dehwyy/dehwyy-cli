package database

import (
	"database/sql"
	"log"
)

type RunnerDB struct {
	db *sql.DB
}

func (r *RunnerDB) Init() {
	r.db = InitDB()
}

func (r *RunnerDB) Close() {
	CloseDB(r.db)
}

func (r *RunnerDB) CreateTableIfNotExists() {
	_, err := r.db.Exec("CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, command TEXT, key TEXT)")
	if err != nil {
		log.Fatalf("Cannot create table: %v", err)
	}
}

func (r* RunnerDB) AddCommandByKey(cmd, key string) {
	_, err := r.db.Exec("INSERT INTO commands (command, key) VALUES (?, ?)", cmd, key)
	if err != nil {
			log.Fatalf("Cannot create template %s with command %s: %v", key, cmd, err)
	}
}

func (r *RunnerDB) DeleteTemplateByKey(key string) int64 {
	result , err := r.db.Exec("DELETE FROM commands WHERE KEY = ?", key)
	if err != nil {
		log.Fatalf("Cannot delete template %s: %v", key, err)
	}

	// Ignoring error due to fact written in docs:
	// Sqlite supports this function
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected
}

func (r* RunnerDB) GetCommandsByKey(key string) *sql.Rows  {
	rows, err := r.db.Query("SELECT command from commands WHERE KEY = ?", key)
	if err != nil {
		log.Fatalf("Cannot query database: %v\n", err)
	}

	return rows
}

func (r *RunnerDB) GetAvailableCommands() *sql.Rows {
	rows, err := r.db.Query("SELECT DISTINCT key FROM commands")
	if err != nil {
		log.Fatalf("Cannot query database: %v\n", err)
	}
	return rows
}

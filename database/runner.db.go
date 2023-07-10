package database

import (
	"database/sql"
	"fmt"

	e "github.com/dehwyy/dehwyy-cli/error-handler"
)

type RunnerDB struct {
	db *sql.DB
}

// Base

func (r *RunnerDB) Init() {
	r.db = InitDB()
}

func (r *RunnerDB) Close() {
	CloseDB(r.db)
}

// Execution

func (r *RunnerDB) CreateTableIfNotExists() {
	 const query = "CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, command TEXT, key TEXT)"

	e.WithFatal(r.db.Exec(query))("Cannot create table")
}

func (r* RunnerDB) AddCommandByKey(cmd, key string) {
	const query = "INSERT INTO commands (command, key) VALUES (?, ?)"

	// message that would appear on error
	errorMessage := fmt.Sprintf("Cannot create template %s with command %s", key, cmd)

	e.WithFatal(r.db.Exec(query, cmd, key))(errorMessage)
}

func (r *RunnerDB) DeleteTemplateByKey(key string) int64 {
	const query = "DELETE FROM commands WHERE KEY = ?"

	// message that would appear on error
	errorMessage := fmt.Sprintf("Cannot delete template %s", key)

	result := e.WithFatal(r.db.Exec(query, key))(errorMessage)
	// Ignoring error due to fact written in docs:
	// Sqlite supports this function
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected
}

// Queries

func (r* RunnerDB) GetCommandsByKey(key string) *sql.Rows  {
	const query = "SELECT command from commands WHERE KEY = ?"

	rows := e.WithFatal(r.db.Query(query, key))("Cannot query database")

	return rows
}

func (r *RunnerDB) GetAvailableCommands() *sql.Rows {
	const query = "SELECT DISTINCT key FROM commands"

	rows := e.WithFatal(r.db.Query(query))("Cannot query database")

	return rows
}

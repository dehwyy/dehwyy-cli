package database

import (
	"database/sql"
	"log"
)

type RunnerDB struct {
	db *sql.DB
}

// must have method for sql interpreter
type SqlDatabase interface {
	New() *sql.DB
}

//

func New(db SqlDatabase) *RunnerDB {
	return &RunnerDB{
		db: db.New (),
	}
}

func (r *RunnerDB) Close() {
	r.db.Close()
}

// Execution

func (r *RunnerDB) CreateTableIfNotExists() {
	 const query = "CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, command TEXT, key TEXT)"

	st, err := r.db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = st.Exec()
	if err != nil {
		log.Fatalf("Cannot create table %v\n", err)
	}
}

func (r* RunnerDB) AddCommandByKey(cmd, key string) {
	const query = "INSERT INTO commands (command, key) VALUES (?, ?)"

	st, err := r.db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = st.Exec(cmd, key)
	if err != nil {
		log.Fatalf("Cannot create template %s with command %s", key, cmd)
	}
}

func (r *RunnerDB) DeleteTemplateByKey(key string) int64 {
	const query = "DELETE FROM commands WHERE KEY = ?"

	st, err := r.db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := st.Exec(key)
	if err != nil {
		log.Fatalf("Cannot delete template %s", key)
	}

	// Ignoring error due to fact written in docs:
	// Sqlite supports this function
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected
}

// Queries

func (r* RunnerDB) GetCommandsByKey(key string) *sql.Rows  {
	const query = "SELECT command from commands WHERE KEY = ?"

	st, err := r.db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := st.Query(key)
	if err != nil {
		log.Fatalf("Cannot query database: %v", err)
	}

	return rows
}

func (r *RunnerDB) GetAvailableCommands() *sql.Rows {
	const query = "SELECT DISTINCT key FROM commands"

	st, err := r.db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := st.Query()
	if err != nil {
		log.Fatalf("Cannot query database: %v", err)
	}

	return rows
}

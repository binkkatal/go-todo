package postgres

import (
	"database/sql"
)

// DB is a PostgreSQL database.
type DB struct {
	*sql.DB
}

// Open opens a new database connection.
func Open(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Begin begins a transaction.
func (db *DB) Begin() *Tx {
	tx, err := db.DB.Begin()
	if err != nil {
		panic(err)
	}
	return &Tx{tx}
}

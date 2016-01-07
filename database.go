package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// DB struct
type DB struct {
	*sqlx.DB
}

// NewDB Create new DB
func NewDB(config *AppConfig) (*DB, error) {
	dbPath := config.DatabasePath
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
		return nil, err
	}
	myDB := DB{db}
	return &myDB, nil
}

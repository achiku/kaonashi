package main

import (
	"log"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB(config *AppConfig) (*DB, error) {
	dbPath := config.DatabasePath
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
		return nil, err
	}
	myDB := DB{db}
	return &myDB, nil
}

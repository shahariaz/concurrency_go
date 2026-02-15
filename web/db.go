package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}

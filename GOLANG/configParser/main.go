package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_ = os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	} else {
		fmt.Println("Successfully connected to the database.")
	}

	// err = CreateTable(db)
	// if err != nil {
	// 	fmt.Println("Error creating table:", err)
	// 	return
	// }

}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT UNIQUE, hashed_password TEXT)")
	return err
}

func InsertUser(db *sql.DB, name string, email string, hashedPassword string) error {
	res, err := db.Exec("INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)", name, email, hashedPassword)
	res.LastInsertId()
	return err
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword string
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	err := db.QueryRow(`SELECT id, name, email, hashed_password FROM users WHERE email = ?`, email).Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	q, err := db.Prepare(`SELECT id, name, email, hashed_password FROM users`)

	if err != nil {
		if err == sql.ErrNoRows {
			return []User{}, nil
		}
		return nil, err
	}

	defer q.Close()
	rows, err := q.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return users, nil

}

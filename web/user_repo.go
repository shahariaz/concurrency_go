package main

import "database/sql"

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword string
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepostory(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (r *UserRepository) CreateTable() error {
	stmt, error := r.DB.Prepare(`
	CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT UNIQUE, hashed_password TEXT)`)
	if error != nil {
		return error
	}
	_, err := stmt.Exec()
	return err
}

func (r *UserRepository) InsertUser(name string, email string, hashedPassword string) error {
	stmt, err := r.DB.Prepare(`INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, email, hashedPassword)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.DB.QueryRow(`SELECT id, name, email, hashed_password FROM users WHERE email = ?`, email).Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

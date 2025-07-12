package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	//"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Username     string
	Email        string
	Password     string
	PasswordHash string
}

func InsertUser(conn *pgx.Conn, username, email, password_hash string) error {
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)",
		username,
		email,
		password_hash,
	)
	if err != nil {
		return fmt.Errorf("Error Insert a User: %w", err)
	}
	return err
}

func GetUsers(conn *pgx.Conn) ([]User, error) {
	rows, err := conn.Query(context.Background(), "SELECT ID, username, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("Error getting users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Email)
		if err != nil {
			return nil, fmt.Errorf("Error to a reading lines: %w", err)
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByEmail(conn *pgx.Conn, email string) (User, error) {
	row := conn.QueryRow(
		context.Background(),
		"SELECT id, username, email, password_hash FROM users WHERE email = $1",
		email,
	)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)

	if err != nil {
		return User{}, fmt.Errorf("Error to reading lines: %w", err)
	}
	return user, err
}

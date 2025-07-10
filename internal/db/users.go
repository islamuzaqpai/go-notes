package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	//"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID           int
	Username     string
	Email        string
	passwordHash string
}

func CreateUser(conn *pgx.Conn, username, email, passwordHash string) error {
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO users (username, emial, passwordHash) VALUES ($1, $2, $3)",
		username,
		email,
		passwordHash,
	)
	if err != nil {
		return fmt.Errorf("Error Insert a User: %w", err)
	}
	return err
}

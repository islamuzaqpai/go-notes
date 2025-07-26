package db

import (
	"context"
	"fmt"
	"github.com/islamuzaqpai/notes-app/internal/config"
	"github.com/jackc/pgx/v5"
	"log"
)

func Connect(cfg *config.Config) *pgx.Conn {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHOST,
		cfg.DBPort,
		cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable conntect to database: %v", err)
	}
	log.Println("Connected to database successfully: %v", conn)
	return conn
}

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"
)

type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

func InsertNote(conn *pgx.Conn, title, content string) error {
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO notes (title, content) VALUES ($1, $2)",
		title,
		content,
	)
	if err != nil {
		return fmt.Errorf("Error Insert a Note: %w", err)
	}
	return nil
}

func GetNotes(conn *pgx.Conn) ([]Note, error) {
	rows, err := conn.Query(context.Background(), "SELECT ID, title, content, created FROM notes")
	if err != nil {
		return nil, fmt.Errorf("Error Insert Note: %w", err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Created)
		if err != nil {
			return nil, fmt.Errorf("Error to reading a lines: %w", err)
		}
		notes = append(notes, n)
	}
	return notes, nil
}

func UpdateNote(conn *pgx.Conn, id int, title, content string) error {
	_, err := conn.Exec(
		context.Background(),
		"UPDATE notes SET title = $1, content = $2 WHERE id = $3",
		title, content, id,
	)
	if err != nil {
		return fmt.Errorf("Error Updating: %w", err)
	}
	return nil
}

func DeleteNote(conn *pgx.Conn, id int) error {
	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM notes WHERE id = $1",
		id,
	)
	if err != nil {
		return fmt.Errorf("Error Deleting: %w", err)
	}
	return nil
}

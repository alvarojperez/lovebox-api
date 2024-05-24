package main

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func openDBConnection() {
	url := getEnvVariable("DATABASE_URL") + "?authToken=" + getEnvVariable("DATABASE_TOKEN")

	var err error
	db, err = sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to open DB %s: %s", url, err)
		os.Exit(1)
	}
}

type Message struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func QueryMessage() (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction: %s", err)
	}

	row := tx.QueryRow("SELECT id, content FROM messages WHERE read_at IS NULL ORDER BY created_at ASC LIMIT 1")

	var message Message

	if err := row.Scan(&message.Id, &message.Content); err != nil {
		log.Printf("Failed to scan message: %s", err)
		tx.Rollback()
		return "", err
	}

	if _, err := tx.Exec("UPDATE messages SET last_retrieved = CURRENT_TIMESTAMP WHERE id = ?", message.Id); err != nil {
		log.Printf("Failed to update message: %s", err)
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Failed to commit transaction: %s", err)
		return "", err
	}

	return message.Content, nil
}

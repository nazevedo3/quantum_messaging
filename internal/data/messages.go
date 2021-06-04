package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Define the struct for the Messages
type Message struct {
	ID      int64
	Message string
	Hash    []byte
}

// Define the MessageModel which holds the DB client
type MessageModel struct {
	DB *sql.DB
}

// Insert interacts with the database to store a message
func (m MessageModel) Insert(message *Message) error {
	query := `
	INSERT INTO messages (message)
	VALUES ($1)
	RETURNING hash
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, message.Message).Scan(&message.Hash)
}

// Get retrieves a message from the database give a hash
func (m MessageModel) Get(hash string) (*Message, error) {
	query := `
	SELECT id, message, hash FROM messages WHERE hash = $1
	`
	var message Message
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Query the database with the provided hash.
	// Store the results in the message variable
	err := m.DB.QueryRowContext(ctx, query, hash).Scan(
		&message.ID,
		&message.Message,
		&message.Hash,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}
	return &message, nil

}

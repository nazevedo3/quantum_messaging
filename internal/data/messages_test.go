package data

import (
	"context"
	"testing"
)

func TestDBInteractions(t *testing.T) {
	ctx := context.Background()

	// Initialize container and the database
	container, db, err := CreateTestContainer(ctx, "testdb")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer container.Terminate(ctx)

	// use the migration files to seed the database
	mig, err := NewMigrator(db)
	if err != nil {
		t.Fatal(err)
	}

	err = mig.Up()
	if err != nil {
		t.Fatal(err)
	}

	// create MessageModel using the test database
	m := &MessageModel{db}

	// create a message that will be stored in the test database
	mockMessage := Message{
		Message: "this is a test",
	}

	err = m.Insert(&mockMessage)
	if err != nil {
		t.Errorf("failed to insert message into the database: %s", err)
	}
	retrieved, err := m.Get(string(mockMessage.Hash))
	if err != nil {
		t.Errorf("failed to get message from the database: %s", err)
	}
	// checks that the hash stored in mockMessage and what was retrieved
	// from the test database are the same
	if string(retrieved.Hash) != string(mockMessage.Hash) {
		t.Errorf("retrieved.Hash (%s) != mockMessage.Hash (%s)", retrieved.Hash, mockMessage.Hash)
	}
	// checks that the message stored in mockMessage and what was retrieved
	// from the test database are the same
	if mockMessage.Message != retrieved.Message {
		t.Errorf("retrieved.Message (%s) != mockMessage.Message (%s)", retrieved.Message, mockMessage.Message)
	}
}

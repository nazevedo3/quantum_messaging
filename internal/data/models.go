package data

import "database/sql"

// Models holds the MessagesModels
type Models struct {
	Messages MessageModel
}

// NewModel intilizes a new Model struct the db client
func NewModel(db *sql.DB) Models {
	return Models{
		Messages: MessageModel{DB: db},
	}
}

package event

import "database/sql"

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	UserID      int
	CreatedAt   string
	UpdatedAt   string
	IsDeleted   bool
	DeletedAt   sql.NullString
}

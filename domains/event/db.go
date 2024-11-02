package event

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

func dbCreateEvent(ctx context.Context, db *sql.Tx, model Event) (int64, error) {
	query := `INSERT INTO events(
	name, description, location, user_id) VALUES(?, ?, ?, ?)`

	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, model.Name, model.Description, model.Location, model.UserID)

	if err != nil {
		slog.Error("error on executing query", "error", err)
		return 0, fmt.Errorf("error executing query: %w", err)
	}

	return result.RowsAffected()
}

func dbGetEvent(ctx context.Context, db *sql.DB) ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

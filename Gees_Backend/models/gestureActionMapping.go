package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

type GestureActionMapping struct {
	ID        *int `json:"id,omitempty"`
	GestureID int  `json:"gesture_id"`
	ActionID  int  `json:"action_id"`
}

// CreateGestureActionMappingTable creates the gesture_action_mapping table if it does not exist
func CreateGestureActionMappingTable(tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS gesture_action_mapping (
		id SERIAL PRIMARY KEY,
		gesture_id INT NOT NULL,
		action_id INT NOT NULL,
		FOREIGN KEY (gesture_id) REFERENCES gestures(id) ON DELETE CASCADE,
		FOREIGN KEY (action_id) REFERENCES actions(id) ON DELETE CASCADE
	)`
	_, err := tx.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create gesture_action_mapping table: %w", err)
	}
	log.Println("gesture_action_mapping table created successfully")
	return nil
}

// ClearGestureActionMappings clears all records from the gesture_action_mapping table
func ClearGestureActionMappings(tx pgx.Tx) error {
	query := `
	DELETE FROM gesture_action_mapping`
	_, err := tx.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to clear gesture_action_mapping: %w", err)
	}
	return nil
}

// InsertGestureActionMapping inserts a new gesture-action mapping into the database
func InsertGestureActionMapping(tx pgx.Tx, mapping *GestureActionMapping) error {
	query := `
	INSERT INTO gesture_action_mapping (gesture_id, action_id)
	VALUES ($1, $2) RETURNING id`
	err := tx.QueryRow(context.Background(), query, mapping.GestureID, mapping.ActionID).Scan(&mapping.ID)
	if err != nil {
		return fmt.Errorf("failed to insert gesture_action_mapping: %w", err)
	}
	return nil
}

// UpdateGestureActionMapping updates an existing gesture-action mapping in the database
func GetGestureActionMappings() ([]GestureActionMapping, error) {
	conn := GetConn()
	if conn == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
	SELECT id, gesture_id, action_id
	FROM gesture_action_mapping`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve gesture_action_mapping: %w", err)
	}
	defer rows.Close()

	var mappings []GestureActionMapping
	for rows.Next() {
		var mapping GestureActionMapping
		if err := rows.Scan(&mapping.ID, &mapping.GestureID, &mapping.ActionID); err != nil {
			return nil, fmt.Errorf("failed to scan gesture_action_mapping row: %w", err)
		}
		mappings = append(mappings, mapping)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", rows.Err())
	}

	return mappings, nil
}

package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

// Action struct
type Action struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name"`
}

// CreateActionTable creates the actions table if it does not exist
func CreateActionTable(tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS actions (
		id SERIAL PRIMARY KEY,
		action TEXT NOT NULL
	)`
	_, err := tx.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("Error creating actions table: %v\n", err)
	}
	log.Printf("Created action table successfully")
	return nil
}

// InsertAction inserts a new action into the actions table
func InsertAction(tx pgx.Tx, action *Action) error {
	if action.ID != nil {
		return fmt.Errorf("ID should be nil when inserting a new action")
	}
	var g = *action
	query := `
	INSERT INTO actions (action)
	VALUES ($1) RETURNING id`
	err := tx.QueryRow(context.Background(), query, g.Name).Scan(&action.ID)
	if err != nil {
		return fmt.Errorf("Error inserting action: %v\n", err)
	}
	log.Printf("Inserted action successfully")
	return nil
}

// UpdateAction updates an existing action in the actions table
func UpdateAction(tx pgx.Tx, action *Action) error {
	if action.ID == nil {
		return fmt.Errorf("action.ID is nil")
	}
	if action.Name == "" {
		return fmt.Errorf("action.Name is empty")
	}

	query := `
	UPDATE actions
	SET action = $1
	WHERE id = $2`
	_, err := tx.Exec(context.Background(), query, action.Name, *action.ID)
	if err != nil {
		return fmt.Errorf("Error updating action: %v\n", err)
	}
	log.Printf("Updated action with ID %d successfully", *action.ID)
	return nil
}

// GetAllActions retrieves all actions from the actions table
func GetAllActions() ([]Action, error) {
	conn := GetConn()
	if conn == nil {
		return nil, fmt.Errorf("Database connection is nil")
	}
	query := `
	SELECT id, action
	FROM actions`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving actions: %v\n", err)
	}
	defer rows.Close()

	var actions []Action
	for rows.Next() {
		var action Action
		err := rows.Scan(&action.ID, &action.Name)
		if err != nil {
			return nil, fmt.Errorf("Error scanning action: %v\n", err)
		}
		actions = append(actions, action)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over actions: %v\n", err)
	}
	return actions, nil
}

// GetActionByID retrieves an action by its ID
func GetActionByID(id int) (*Action, error) {
	conn := GetConn()
	if conn == nil {
		return nil, fmt.Errorf("Database connection is nil")
	}

	var action Action
	query := `
    SELECT id, action
    FROM actions
    WHERE id = $1`

	err := conn.QueryRow(context.Background(), query, id).Scan(&action.ID, &action.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("Error retrieving action: %v\n", err)
	}

	return &action, nil
}

package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

type Gesture struct {
	ID          *int    `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	ImageUrl    *string `json:"image_url,omitempty"`
}

// CreateGestureTable creates the gestures table if it does not exist
func CreateGestureTable(tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS gestures (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		image_url TEXT
	)`
	_, err := tx.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("Error creating gestures table: %v\n", err)
	}
	log.Printf("Created gesture table successfully")
	return nil
}

// InsertGesture inserts a new gesture into the gestures table
func InsertGesture(tx pgx.Tx, gesture *Gesture) error {
	if gesture.Name == "" {
		return fmt.Errorf("gesture.Name is empty")
	}
	if gesture.ID != nil {
		return fmt.Errorf("gesture.ID is not nil")
	}
	query := `
	INSERT INTO gestures (name, description, image_url)
	VALUES ($1, $2, $3) RETURNING id`
	err := tx.QueryRow(context.Background(), query, gesture.Name, gesture.Description, gesture.ImageUrl).Scan(&gesture.ID)
	if err != nil {
		return fmt.Errorf("Error inserting gesture: %v\n", err)
	}
	log.Printf("Inserted gesture successfully")
	return nil
}

// UpdateGesture updates an existing gesture in the gestures table
func UpdateGesture(tx pgx.Tx, gesture *Gesture) error {
	if gesture.ID == nil {
		return fmt.Errorf("gesture.ID is nil")
	}
	if gesture.Name == "" {
		return fmt.Errorf("gesture.Name is empty")
	}

	query := `
	UPDATE gestures
	SET name = $1, description = $2, image_url = $3
	WHERE id = $4`
	_, err := tx.Exec(context.Background(), query, gesture.Name, gesture.Description, gesture.ImageUrl, *gesture.ID)
	if err != nil {
		return fmt.Errorf("Error updating gesture: %v\n", err)
	}
	log.Printf("Updated gesture successfully")
	return nil
}

// GetAllGestures retrieves all gestures from the gestures table
func GetAllGestures() ([]Gesture, error) {
	conn := GetConn()
	if conn == nil {
		return nil, fmt.Errorf("Database connection is nil")
	}
	query := `
	SELECT id, name, description, image_url
	FROM gestures`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving gestures: %v\n", err)
	}
	defer rows.Close()

	var gestures []Gesture
	for rows.Next() {
		var gesture Gesture
		err := rows.Scan(&gesture.ID, &gesture.Name, &gesture.Description, &gesture.ImageUrl)
		if err != nil {
			return nil, fmt.Errorf("Error scanning gesture: %v\n", err)
		}
		gestures = append(gestures, gesture)
	}
	return gestures, nil
}

// GetGestureByID retrieves a gesture by its ID from the gestures table
func GetGestureByID(id int) (*Gesture, error) {
	conn := GetConn()
	if conn == nil {
		return nil, fmt.Errorf("Database connection is nil")
	}
	query := `
	SELECT id, name, description, image_url
	FROM gestures WHERE id = $1`
	var gesture Gesture
	err := conn.QueryRow(context.Background(), query, id).Scan(&gesture.ID, &gesture.Name, &gesture.Description, &gesture.ImageUrl)
	if err != nil {
		return &gesture, fmt.Errorf("Error retrieving gesture: %v\n", err)
	}
	return &gesture, nil
}

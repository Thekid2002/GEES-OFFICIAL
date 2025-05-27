package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

var conn *pgx.Conn

func Connect() error {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Connected to the database successfully")
	return nil
}

func GetConn() *pgx.Conn {
	if conn == nil {
		err := Connect()
		if err != nil {
			log.Fatalf("Error connecting to database: %v\n", err)
		}
	}
	return conn
}

// StartTransaction starts a new database transaction
func StartTransaction() (pgx.Tx, error) {
	conn := GetConn()
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Error starting transaction: %v\n", err)
	}
	return tx, nil
}

// CommitTransaction commits the given transaction
func CommitTransaction(tx pgx.Tx) error {
	err := tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("Error committing transaction: %v\n", err)
	}
	return nil
}

// RollbackTransaction rolls back the given transaction
func RollbackTransaction(tx pgx.Tx) error {
	err := tx.Rollback(context.Background())
	if err != nil {
		return fmt.Errorf("Error rolling back transaction: %v\n", err)
	}
	return nil
}

func InitDatabases() error {
	tx, err := StartTransaction()
	if err != nil {
		return fmt.Errorf("Error starting transaction: %v\n", err)
	}
	err = CreateGestureTable(tx)
	if err != nil {
		rollErr := RollbackTransaction(tx)
		if rollErr != nil {
			log.Fatalf("Error rolling back transaction: %v\n", rollErr)
		}
		return fmt.Errorf("Error creating Gesture table: %v\n", err)
	}

	err = CreateFeatureTable(tx)
	if err != nil {
		rollErr := RollbackTransaction(tx)
		if rollErr != nil {
			log.Fatalf("Error rolling back transaction: %v\n", rollErr)
		}
		return fmt.Errorf("Error creating Feature table: %v\n", err)
	}
	err = CreateDataPointTable(tx)
	if err != nil {
		rollErr := RollbackTransaction(tx)
		if rollErr != nil {
			log.Fatalf("Error rolling back transaction: %v\n", rollErr)
		}
		return fmt.Errorf("Error creating DataPoint table: %v\n", err)
	}

	err = CreateActionTable(tx)
	if err != nil {
		rollErr := RollbackTransaction(tx)
		if rollErr != nil {
			log.Fatalf("Error rolling back transaction: %v\n", rollErr)
		}
		return fmt.Errorf("Error creating Action table: %v\n", err)
	}

	err = CreateGestureActionMappingTable(tx)
	if err != nil {
		rollErr := RollbackTransaction(tx)
		if rollErr != nil {
			log.Fatalf("Error rolling back transaction: %v\n", rollErr)
		}
		return fmt.Errorf("Error creating GestureActionMapping table: %v\n", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		return fmt.Errorf("Error committing transaction: %v\n", err)
	}

	fmt.Printf("Created database successfully \n")

	return nil
}

func InitializeDataIntoDatabase() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("Error connecting to database: %v\n", err)
	}
	fmt.Println("Connected to the database successfully")

	tx, err := StartTransaction()
	if err != nil {
		return fmt.Errorf("Error starting transaction: %v\n", err)
	}

	ges1 := "A waving gesture"
	ges2 := "A thumbs-up gesture"
	ges3 := "A clapping gesture"
	ges4 := "A pointing gesture"
	ges5 := "A fist gesture"

	pic1 := "/tutorial_for_mapping/tutorial1.png"
	pic2 := "/tutorial_for_mapping/tutorial2.png"
	pic3 := "/tutorial_for_mapping/tutorial3.png"
	pic4 := "/tutorial_for_mapping/tutorial4.png"
	pic5 := "/tutorial_for_mapping/tutorial5.png"

	id1 := 1
	id2 := 2
	id3 := 3
	id4 := 4
	id5 := 5

	gestures := []Gesture{
		{ID: &id1, Name: "Wave", Description: &ges1, ImageUrl: &pic1},
		{ID: &id2, Name: "Thumbs Up", Description: &ges2, ImageUrl: &pic2},
		{ID: &id3, Name: "Clap", Description: &ges3, ImageUrl: &pic3},
		{ID: &id4, Name: "Point", Description: &ges4, ImageUrl: &pic4},
		{ID: &id5, Name: "Fist", Description: &ges5, ImageUrl: &pic5},
	}

	for i := range gestures {
		err = InsertGesture(tx, &gestures[i])
		if err != nil {
			err := RollbackTransaction(tx)
			if err != nil {
				return err
			}
			return fmt.Errorf("Error inserting gesture: %v\n", err)
		}
	}

	id1 = 1
	id2 = 2
	id3 = 3
	id4 = 4
	id5 = 5
	id6 := 6

	// Insert actions
	actions := []Action{
		{ID: &id1, Name: "Open window"},
		{ID: &id2, Name: "Close window"},
		{ID: &id3, Name: "Increase thermostat"},
		{ID: &id4, Name: "Decrease thermostat"},
		{ID: &id5, Name: "Turn on lights"},
		{ID: &id6, Name: "Turn off lights"},
	}
	for i := range actions {
		err = InsertAction(tx, &actions[i])
		if err != nil {
			err := RollbackTransaction(tx)
			if err != nil {
				return err
			}
			return fmt.Errorf("Error inserting action: %v\n", err)
		}
	}

	err = CommitTransaction(tx)
	if err != nil {
		return fmt.Errorf("Error committing transaction: %v\n", err)
	}

	fmt.Println("Initial data inserted successfully")
	return nil
}

func Close() error {
	err := conn.Close(context.Background())
	if err != nil {
		return fmt.Errorf("Error closing database connection: %v\n", err)
	}
	return nil
}

func ConnectTest() error {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("Unable to connect to test database: %v\n", err)
	}
	fmt.Println("Connected to the test database successfully")
	return nil
}

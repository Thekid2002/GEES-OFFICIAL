package models

import (
	"Gees_Backend/env"
	"context"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Runs before the testing
	envString := "test.env"
	env.LoadEnv(&envString)
	err := Connect()
	if err != nil {
		panic(err.Error())
	}

	err = InitDatabases()
	if err != nil {
		log.Fatalf("Failed to initialize databases: %v", err)
	}

	code := m.Run()

	tx, err := StartTransaction()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}

	_, err = tx.Exec(context.Background(), `DROP TABLE IF EXISTS gesture_action_mapping CASCADE; DROP TABLE IF EXISTS feature_data CASCADE; DROP TABLE IF EXISTS gestures CASCADE; DROP TABLE IF EXISTS data_points CASCADE; DROP TABLE IF EXISTS actions CASCADE`)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	defer func() {
		if err != nil {
			err = RollbackTransaction(tx)
			if err != nil {
				log.Fatalf("Failed to rollback transaction: %v", err)
			}
		}
	}()

	err = CommitTransaction(tx)
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
	log.Println("Successfully cleaned up database tables")
	os.Exit(code)
}

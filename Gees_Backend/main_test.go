package main

import (
	"Gees_Backend/env"
	"Gees_Backend/models"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	envString := "test.env"
	env.LoadEnv(&envString)
	err := models.Connect()
	if err != nil {
		panic("Failed to connect to the database")
	}
	code := m.Run()

	tx, err := models.StartTransaction()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}

	_, err = tx.Exec(context.Background(), `DROP TABLE IF EXISTS gesture_action_mapping CASCADE; DROP TABLE IF EXISTS feature_data CASCADE; DROP TABLE IF EXISTS gestures CASCADE; DROP TABLE IF EXISTS data_points CASCADE; DROP TABLE IF EXISTS actions CASCADE`)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	defer func() {
		if err != nil {
			err = models.RollbackTransaction(tx)
			if err != nil {
				log.Fatalf("Failed to rollback transaction: %v", err)
			}
		}
	}()

	err = models.CommitTransaction(tx)
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
	os.Exit(code)
}

func TestInitFunction(t *testing.T) {
	initialize()
	t.Log("Init function test executed successfully.")
}

func TestCorsMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := &httptest.ResponseRecorder{}

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := corsMiddleware(dummyHandler)

	handler.ServeHTTP(rr, req)

	if rr.Header().Get("Access-Control-Allow-Methods") != "GET, POST, PUT, DELETE, OPTIONS" {
		t.Errorf("Expected CORS methods header, got: %v", rr.Header().Get("Access-Control-Allow-Methods"))
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got: %v", rr.Code)
	}
}

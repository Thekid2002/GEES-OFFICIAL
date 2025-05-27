package controllers

import (
	"Gees_Backend/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestEditActionNoBody(t *testing.T) {
	req, err := http.NewRequest("PUT", "/edit-action", nil)
	if err != nil {
		t.Fatal(err)
	}

	resW := httptest.NewRecorder()

	EditAction(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
	}
}

func TestEditActionInvalidJSON(t *testing.T) {
	// ActionName should be called Name, thus invalid JSON
	body := strings.NewReader(`{"ID": 1, "ActionName": "Test Action"}`)
	req, err := http.NewRequest("PUT", "/edit-action", body)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	EditAction(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
	}
}

// TestEditActionNoID tests the case where the action ID is not provided
func TestEditActionValid(t *testing.T) {
	tx, err := models.StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}
	action := models.Action{Name: "Test Action"}

	err = models.InsertAction(tx, &action)

	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}
	err = models.CommitTransaction(tx)

	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	body := strings.NewReader(`{"ID": ` + strconv.Itoa(*action.ID) + `, "Name": "New action name"}`)
	req, err := http.NewRequest("PUT", "/edit-action", body)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	EditAction(resW, req)

	if resW.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resW.Code)
		return
	}

	updatedAction, err := models.GetActionByID(*action.ID)
	if err != nil {
		t.Fatalf("Failed to get action by ID: %v", err)
	}
	if updatedAction.Name != "New action name" {
		t.Errorf("Expected action name 'New action name', got '%s'", updatedAction.Name)
	}
}

func FuzzCreateAction(f *testing.F) {
	f.Add(`{"ID": 1, "Name": "Test Action"}`)
	f.Fuzz(func(t *testing.T, body string) {
		req, err := http.NewRequest("POST", "/create-action", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		resW := httptest.NewRecorder()

		CreateAction(resW, req)

		if resW.Code != http.StatusOK && resW.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d or %d, got %d", http.StatusOK, http.StatusBadRequest, resW.Code)
		}
	})
}

func TestCreateActionInvalidJSON(t *testing.T) {
	body := strings.NewReader(`{"ID": 1, "ActionName": "Test Action"}`)
	req, err := http.NewRequest("POST", "/create-action", body)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	CreateAction(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
	}
}

func TestCreateActionWithID(t *testing.T) {
	body := strings.NewReader(`{"id": "10", "name": "Test Action"}`)
	req, err := http.NewRequest("POST", "/create-action", body)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	CreateAction(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
	}
}

func TestCreateActionValid(t *testing.T) {
	body := strings.NewReader(`{"name": "Test Action"}`)
	req, err := http.NewRequest("POST", "/create-action", body)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	CreateAction(resW, req)

	if resW.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resW.Code)
		return
	}

	var action models.Action
	err = json.Unmarshal(resW.Body.Bytes(), &action)

	if action.Name != "Test Action" {
		t.Errorf("Expected action name 'Test Action', got '%s'", action.Name)
	}
}

func TestCreateActionEmptyName(t *testing.T) {
	body := strings.NewReader(`{"name": ""}`)
	req, err := http.NewRequest("POST", "/create-action", body)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	CreateAction(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
		return
	}

	var action models.Action
	err = json.Unmarshal(resW.Body.Bytes(), &action)

	if action.Name != "" {
		t.Errorf("Expected action name '', got '%s'", action.Name)
	}
}

func TestGetActionByID(t *testing.T) {
	tx, err := models.StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}
	action := models.Action{Name: "Test Action"}
	err = models.InsertAction(tx, &action)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}
	err = models.CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	req, err := http.NewRequest("GET", "/action/"+strconv.Itoa(*action.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	GetActionByID(resW, req)

	if resW.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resW.Code)
		return
	}

	var retrievedAction models.Action
	err = json.Unmarshal(resW.Body.Bytes(), &retrievedAction)

	if retrievedAction.Name != "Test Action" {
		t.Errorf("Expected action name 'Test Action', got '%s'", retrievedAction.Name)
	}
}

func TestGetActionByIDInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/action/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	GetActionByID(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
	}
}

func TestGetActions(t *testing.T) {
	tx, err := models.StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}
	action := models.Action{Name: "Test Action"}
	err = models.InsertAction(tx, &action)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}
	err = models.CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	req, err := http.NewRequest("GET", "/get-actions", nil)
	if err != nil {
		t.Fatal(err)
	}
	resW := httptest.NewRecorder()

	GetActions(resW, req)

	if resW.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resW.Code)
		return
	}

	var actions []models.Action
	err = json.Unmarshal(resW.Body.Bytes(), &actions)

	if len(actions) == 0 {
		t.Errorf("Expected at least one action, got none")
	}
}

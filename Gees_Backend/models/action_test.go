package models

import "testing"

func TestInsertAction(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	action := Action{
		Name: "Test Action",
	}

	err = InsertAction(tx, &action)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}

func TestUpdateAction(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	action := &Action{
		Name: "Test Action",
	}

	err = InsertAction(tx, action)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}

	action.Name = "Updated Action"
	err = UpdateAction(tx, action)
	if err != nil {
		t.Fatalf("Failed to update action: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	updatedAction, err := GetActionByID(*action.ID)
	if err != nil {
		t.Fatalf("Failed to get action by ID: %v", err)
	}

	if updatedAction.Name != "Updated Action" {
		t.Errorf("Expected action name 'Updated Action', got '%s'", action.Name)
	}
}

func TestGetActionByID(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	action := &Action{
		Name: "Test Action",
	}

	err = InsertAction(tx, action)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	retrievedAction, err := GetActionByID(*action.ID)
	if err != nil {
		t.Fatalf("Failed to get action by ID: %v", err)
	}

	if retrievedAction.Name != "Test Action" {
		t.Errorf("Expected action name 'Test Action', got '%s'", retrievedAction.Name)
	}
}

func TestGetAllActions(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	action1 := &Action{
		Name: "Test Action 1",
	}
	action2 := &Action{
		Name: "Test Action 2",
	}

	err = InsertAction(tx, action1)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}

	err = InsertAction(tx, action2)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	actions, err := GetAllActions()
	if err != nil {
		t.Fatalf("Failed to get all actions: %v", err)
	}

	if len(actions) < 2 {
		t.Errorf("Expected at least 2 actions, got %d", len(actions))
	}
}

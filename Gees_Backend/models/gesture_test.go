package models

import "testing"

func TestInsertGesture(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	gesture := &Gesture{
		Name: "Test Gesture",
	}

	err = InsertGesture(tx, gesture)
	if err != nil {
		t.Fatalf("Failed to insert gesture: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	gesture, err = GetGestureByID(*gesture.ID)
	if err != nil {
		t.Fatalf("Failed to get gesture by ID: %v", err)
	}

	if gesture.Name != "Test Gesture" {
		t.Errorf("Expected gesture name 'Test Gesture', got '%s'", gesture.Name)
	}
}

func TestInsertGestureWithNoName(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	gesture := Gesture{
		Name: "",
	}

	err = InsertGesture(tx, &gesture)
	if err == nil {
		t.Fatalf("Expected error when inserting gesture with no name, but got none")
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}

func TestInsertGestureWithID(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	id1 := 1
	gesture := Gesture{
		ID:   &id1,
		Name: "Test Gesture with ID",
	}

	err = InsertGesture(tx, &gesture)
	if err == nil {
		t.Fatalf("Failed to fail an insert of gesture with ID: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}

func TestGetAllGestures(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	gesture := Gesture{
		Name: "Test Gesture",
	}

	err = InsertGesture(tx, &gesture)
	if err != nil {
		t.Fatalf("Failed to insert gesture: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
	gestures, err := GetAllGestures()
	if err != nil {
		t.Fatalf("Failed to get all gestures: %v", err)
	}

	if len(gestures) == 0 {
		t.Errorf("Expected at least one gesture, got %d", len(gestures))
	}
}

func TestGetGestureByNoID(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	gesture := Gesture{
		Name: "Test Gesture",
	}

	err = InsertGesture(tx, &gesture)
	if err != nil {
		t.Fatalf("Failed to insert gesture: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	_, err = GetGestureByID(0)
	if err == nil {
		t.Fatalf("Expected error when getting gesture by ID 0, but got none")
	}
}

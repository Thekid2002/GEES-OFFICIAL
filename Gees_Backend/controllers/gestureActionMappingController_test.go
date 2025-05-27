package controllers

import (
	"Gees_Backend/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateGestureActionMappings(t *testing.T) {
	gesture1 := models.Gesture{Name: "Test Gesture 1"}
	action1 := models.Action{Name: "Test Action 1"}
	gesture2 := models.Gesture{Name: "Test Gesture 2"}
	action2 := models.Action{Name: "Test Action 2"}

	tx, err := models.StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	err = models.InsertGesture(tx, &gesture1)
	if err != nil {
		t.Fatalf("Failed to insert gesture: %v", err)
	}
	err = models.InsertAction(tx, &action1)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}
	err = models.InsertGesture(tx, &gesture2)
	if err != nil {
		t.Fatalf("Failed to insert gesture: %v", err)
	}
	err = models.InsertAction(tx, &action2)
	if err != nil {
		t.Fatalf("Failed to insert action: %v", err)
	}
	err = models.CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	body := strings.NewReader(fmt.Sprintf(`[{"gesture_id": %d, "action_id": %d}, {"gesture_id": %d, "action_id": %d}]`,
		*gesture1.ID, *action1.ID, *gesture2.ID, *action2.ID))
	req, err := http.NewRequest("PUT", "/update-gesture-action-mappings", body)

	if err != nil {
		t.Fatal(err)
	}

	resW := httptest.NewRecorder()

	UpdateGestureActionMappings(resW, req)

	if resW.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resW.Code)
		return
	}

	mappings, err := models.GetGestureActionMappings()

	if err != nil {
		t.Errorf("Error retrieving gesture action mappings: %v", err)
		return
	}

	if len(mappings) < 2 {
		t.Errorf("Expected atleast 2 mappings, got %d", len(mappings))
		return
	}
}

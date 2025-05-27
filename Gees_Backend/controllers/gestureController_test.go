package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGestures(t *testing.T) {
	req, err := http.NewRequest("GET", "/get-gestures", nil)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resW := httptest.NewRecorder()

	GetGestures(resW, req)

	if resW.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resW.Code)
		return
	}
}

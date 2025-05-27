package controllers

import (
	"Gees_Backend/models"
	"encoding/json"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func randomFloat() float64 {
	return rand.Float64() * 100
}

func fillStructWithRandomValues(s interface{}) {
	rand.Seed(time.Now().UnixNano())

	v := reflect.ValueOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.CanSet() {
			if field.Kind() == reflect.Float64 {
				field.SetFloat(randomFloat())
			} else if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Float64 {
				ptr := randomFloat()
				field.Set(reflect.ValueOf(&ptr))
			}
		}
	}
}

func compareFeatureData(retrieved, expected models.FeatureData) bool {
	return *retrieved.AccMeanX == *expected.AccMeanX &&
		*retrieved.AccMeanY == *expected.AccMeanY &&
		*retrieved.AccMeanZ == *expected.AccMeanZ &&
		*retrieved.AccVarianceX == *expected.AccVarianceX &&
		*retrieved.AccVarianceY == *expected.AccVarianceY &&
		*retrieved.AccVarianceZ == *expected.AccVarianceZ &&
		*retrieved.AccMedianX == *expected.AccMedianX &&
		*retrieved.AccMedianY == *expected.AccMedianY &&
		*retrieved.AccMedianZ == *expected.AccMedianZ &&
		*retrieved.AccStdDevX == *expected.AccStdDevX &&
		*retrieved.AccStdDevY == *expected.AccStdDevY &&
		*retrieved.AccStdDevZ == *expected.AccStdDevZ &&
		*retrieved.GyrMeanX == *expected.GyrMeanX &&
		*retrieved.GyrMeanY == *expected.GyrMeanY &&
		*retrieved.GyrMeanZ == *expected.GyrMeanZ
}

func TestWebSocketMocking(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade WebSocket: %v", err)
		}
		defer conn.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				break
			}
			conn.WriteMessage(messageType, message)
		}
	}))
	defer server.Close()

	// Create a WebSocket client connection
	wsURL := "ws" + server.URL[len("http"):]
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()

	GlobalGestureWS = conn

	err = GlobalGestureWS.WriteMessage(websocket.TextMessage, []byte("test message"))
	if err != nil {
		t.Fatalf("Failed to write message: %v", err)
	}

	_, message, err := GlobalGestureWS.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	if string(message) != "test message" {
		t.Errorf("Expected 'test message', got '%s'", message)
	}
}

func TestRecordArduinoFeatureDataWithNoBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/feature-data", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resW := httptest.NewRecorder()

	recordArduinoFeatureData(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
		return
	}
}

func TestRecordArduinoFeatureDataWithInvalidBody(t *testing.T) {
	body := strings.NewReader(`{"ID": 1, "FeatureName": "Test Feature", "FeatureValue": "Test Value"}`)
	req, err := http.NewRequest("POST", "/feature-data", body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	tx, err := models.StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}
	gesture := &models.Gesture{
		Name: "Test Gesture",
	}

	err = models.InsertGesture(tx, gesture)
	if err != nil {
		t.Fatalf("Failed to insert gesture: %v", err)
	}
	err = models.CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	GlobalReadyForRecordingData = true
	GlobalActiveGestureID = *gesture.ID
	resW := httptest.NewRecorder()

	recordArduinoFeatureData(resW, req)

	if resW.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resW.Code)
		return
	}
}

func TestRecordArduinoFeatureData(t *testing.T) {
	featureData := models.FeatureData{}
	fillStructWithRandomValues(&featureData)
	featureDataJson, err := json.Marshal(featureData)
	body := strings.NewReader(string(featureDataJson))
	req, err := http.NewRequest("POST", "/feature-data", body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	tx, err := models.StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}
	gesture := &models.Gesture{
		Name: "Test Gesture",
	}

	err = models.InsertGesture(tx, gesture)
	if err != nil {
		t.Fatalf("Failed to insert gesture: %v", err)
	}
	err = models.CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	GlobalReadyForRecordingData = true
	GlobalActiveGestureID = *gesture.ID
	resW := httptest.NewRecorder()

	recordArduinoFeatureData(resW, req)

	if resW.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resW.Code)
		return
	}
}

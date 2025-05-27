package controllers

import (
	"Gees_Backend/models"
	"Gees_Backend/services"
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

var GlobalActiveGestureID = -1           // Track active gesture state
var GlobalReadyForRecordingData = false  // Tell arduino to make IMU measurements
var GlobalReadyForValidatingData = false // Tell arduino to make IMU measurements
var GlobalGestureWS *websocket.Conn      // Websocket to the frontend gesture tracking

func featureDataFailure(err error) {
	GlobalReadyForRecordingData = false
	GlobalReadyForValidatingData = false
	log.Printf("Feature Data Error: %s", err)
}

// ArduinoPostFeatureData is an endpoint to receive and store feature data from the Arduino
func ArduinoPostFeatureData(resW http.ResponseWriter, req *http.Request) {
	if GlobalReadyForRecordingData {
		recordArduinoFeatureData(resW, req)
		return
	}

	if GlobalReadyForValidatingData {
		validateArduinoFeatureData(resW, req)
		return
	}

	http.Error(resW, "Not ready to record data", http.StatusBadRequest)
	return
}

// recordArduinoFeatureData is a helper function to handle the recording of feature data
func recordArduinoFeatureData(resW http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(resW, "Request body is nil", http.StatusBadRequest)
		log.Println("Request body is nil")
		return
	}
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(resW, "Invalid data", http.StatusBadRequest)
		log.Printf("Error reading request body: %s", err)
		return
	}

	var featureData models.FeatureData
	err = json.Unmarshal(body, &featureData)
	if err != nil {
		http.Error(resW, "Error parsing featuredata", http.StatusBadRequest)
		log.Printf("Error parsing featuredata: %s", err)
		return
	}

	err = models.ValidateFeatureDataNotNil(featureData)
	if err != nil {
		http.Error(resW, "Feature data is nil", http.StatusBadRequest)
		log.Printf("Feature data is nil: %s", err)
		return
	}

	msg := WSMessage{
		MsgType: STOP,
		Data:    models.ConvertDataPointsToJSON(featureData.DataPoints),
	}
	err = WebsocketStopCountdown(GlobalGestureWS, &msg)

	if err != nil {
		http.Error(resW, "Error stopping countdown", http.StatusInternalServerError)
		log.Printf("Error stopping countdown: %v", err)
		return
	}

	if GlobalActiveGestureID == -1 {
		http.Error(resW, "No gesture was assigned, you fool!", http.StatusInternalServerError)
		log.Printf("No gesture was assigned, you fool!")
		GlobalReadyForRecordingData = false
		return
	}

	tx, err := models.StartTransaction()
	if err != nil {
		http.Error(resW, "Error starting transaction", http.StatusInternalServerError)
		featureDataFailure(err)
		return
	}

	defer func() {
		if err != nil {
			models.RollbackTransaction(tx)
		}
	}()

	err = models.InsertFeatureData(tx, &featureData, GlobalActiveGestureID)
	if err != nil {
		http.Error(resW, "Feature data was not inserted correctly!", http.StatusInternalServerError)
		featureDataFailure(err)
		return
	}

	for _, dataPoint := range featureData.DataPoints {
		dataPoint.FeatureDataID = featureData.ID
		dp := dataPoint
		err = models.InsertDataPoint(tx, &dp)
		if err != nil {
			http.Error(resW, "Error storing data", http.StatusInternalServerError)
			featureDataFailure(err)
			return
		}
	}

	if err := models.CommitTransaction(tx); err != nil {
		http.Error(resW, "Error committing transaction", http.StatusInternalServerError)
		featureDataFailure(err)
		return
	}

	WebsocketStopCountdown(GlobalGestureWS, &msg)
	resW.WriteHeader(http.StatusOK)
	json.NewEncoder(resW).Encode(map[string]string{"message": "Feature data received and stored"})
	return
}

// validateArduinoFeatureData is a helper function to handle the validation of feature data
func validateArduinoFeatureData(resW http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(resW, "Invalid data", http.StatusBadRequest)
		log.Printf("Error reading request body: %s", err)
		return
	}

	var featureData models.FeatureData
	err = json.Unmarshal(body, &featureData)
	if err != nil {
		http.Error(resW, "Error parsing featuredata", http.StatusBadRequest)
		log.Printf("Error parsing featuredata: %s", err)
		return
	}
	err = models.ValidateFeatureDataNotNil(featureData)

	if err != nil {
		http.Error(resW, "Feature data is nil", http.StatusBadRequest)
		log.Printf("Feature data is nil: %s", err)
		return
	}

	msg := WSMessage{
		MsgType: STOP,
		Data:    models.ConvertDataPointsToJSON(featureData.DataPoints),
	}

	predictionResponse, err := services.PredictGestureFromCompleteFeatureData(featureData)

	if err != nil {
		log.Printf("Error sending data to prediction service: %v", err)
		http.Error(resW, "Error sending data to prediction service", http.StatusInternalServerError)
		return
	}
	predictionResponseJson, err := json.Marshal(predictionResponse)

	if err != nil {
		log.Printf("Error marshaling prediction response: %v", err)
		http.Error(resW, "Error marshaling prediction response", http.StatusInternalServerError)
		return
	}

	msg = WSMessage{
		MsgType: PREDICTION,
		Data:    string(predictionResponseJson),
	}

	dataStr, ok := msg.Data.(string)
	if !ok {
		log.Printf("Error: msg.Data is not a string")
		http.Error(resW, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = GlobalGestureWS.WriteMessage(websocket.TextMessage, []byte(dataStr))
	if err != nil {
		log.Printf("Error writing message to WebSocket: %v", err)
		return
	}
}

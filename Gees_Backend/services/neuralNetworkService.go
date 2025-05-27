package services

import (
	"Gees_Backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type PropagateResult struct {
	Prediction *float64 `json:"prediction,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	TimeStamp  string   `json:"timeStamp,omitempty"`
	Error      string   `json:"error,omitempty"`
}

// PredictGestureFromCompleteFeatureData sends feature data to the Python app for prediction
func PredictGestureFromCompleteFeatureData(featureData models.FeatureData) (*PropagateResult, error) {
	body, err := json.Marshal(featureData)
	if err != nil {
		log.Printf("Error marshaling feature data: %v", err)
		return nil, err
	}

	res, err := http.Post(os.Getenv("PYTHON_URL")+"/predict", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error sending data to prediction service: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	var result PropagateResult
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return nil, err
	}

	if result.Error != "" {
		log.Printf("Error from prediction service: %s", result.Error)
		return nil, fmt.Errorf("prediction service error: %s", result.Error)
	}

	return &result, nil
}

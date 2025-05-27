package controllers

import (
	"Gees_Backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetGestures an endpoint to get all gestures
func GetGestures(resW http.ResponseWriter, req *http.Request) {
	tx, err := models.StartTransaction()
	if err != nil {
		http.Error(resW, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			log.Printf("Error rolling back transaction: %v", err)
			if rollbackErr := models.RollbackTransaction(tx); rollbackErr != nil {
				panic(rollbackErr)
			}
		}
	}()

	gestures, err := models.GetAllGestures()
	if err != nil {
		log.Printf("Error retrieving gestures: %v", err)
		http.Error(resW, "Error retrieving gestures", http.StatusInternalServerError)
		return
	}

	resW.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(resW).Encode(gestures)
	if err != nil {
		http.Error(resW, err.Error(), http.StatusInternalServerError)
	}
}

// GetGestureByID an endpoint to get a gesture by ID
func GetGestureByID(resW http.ResponseWriter, req *http.Request) {
	gestureIDString := req.URL.Path
	gestureIDString = gestureIDString[len("/gesture/"):]
	if gestureIDString == "" {
		log.Println("Missing gesture ID")
		http.Error(resW, "Missing gesture ID", http.StatusBadRequest)
		return
	}
	gestureID, err := strconv.Atoi(gestureIDString)
	if err != nil {
		log.Printf("Error converting gesture ID to int: %v", err)
		http.Error(resW, "Invalid gesture ID", http.StatusBadRequest)
		return
	}

	gesture, err := models.GetGestureByID(gestureID)
	if err != nil {
		log.Printf("Error retrieving gesture: %v", err)
		http.Error(resW, "Error retrieving gesture", http.StatusInternalServerError)
		return
	}

	resW.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(resW).Encode(gesture)
	if err != nil {
		log.Printf("Error encoding gesture to JSON: %v", err)
		http.Error(resW, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateGesture an endpoint to create a new gesture
func CreateGesture(resW http.ResponseWriter, req *http.Request) {
	var gesture models.Gesture
	err := json.NewDecoder(req.Body).Decode(&gesture)
	if err != nil {
		log.Printf("Error decoding gesture: %v", err)
		http.Error(resW, "Error decoding gesture", http.StatusBadRequest)
		return
	}

	tx, err := models.StartTransaction()
	if err != nil {
		http.Error(resW, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			log.Printf("Error rolling back transaction: %v", err)
			if rollbackErr := models.RollbackTransaction(tx); rollbackErr != nil {
				panic(rollbackErr)
			}
		}
	}()

	err = models.InsertGesture(tx, &gesture)
	if err != nil {
		log.Printf("Error inserting gesture: %v", err)
		http.Error(resW, "Error inserting gesture", http.StatusInternalServerError)
		return
	}

	err = models.CommitTransaction(tx)
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		http.Error(resW, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	resW.WriteHeader(http.StatusCreated)
}

// EditGesture an endpoint to update an existing gesture
func EditGesture(resW http.ResponseWriter, req *http.Request) {
	var gesture models.Gesture
	err := json.NewDecoder(req.Body).Decode(&gesture)
	if err != nil {
		log.Printf("Error decoding gesture: %v", err)
		http.Error(resW, "Error decoding gesture", http.StatusBadRequest)
		return
	}

	tx, err := models.StartTransaction()
	if err != nil {
		http.Error(resW, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			log.Printf("Error rolling back transaction: %v", err)
			if rollbackErr := models.RollbackTransaction(tx); rollbackErr != nil {
				panic(rollbackErr)
			}
		}
	}()

	err = models.UpdateGesture(tx, &gesture)
	if err != nil {
		log.Printf("Error updating gesture: %v", err)
		http.Error(resW, "Error updating gesture", http.StatusInternalServerError)
		return
	}

	err = models.CommitTransaction(tx)
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		http.Error(resW, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	resW.WriteHeader(http.StatusOK)
}

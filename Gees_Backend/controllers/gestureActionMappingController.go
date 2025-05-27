package controllers

import (
	"Gees_Backend/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// UpdateGestureActionMappings is an endpoint to update gesture action mappings
func UpdateGestureActionMappings(resW http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(resW, "Invalid data", http.StatusBadRequest)
		return
	}

	var gestureMappings []models.GestureActionMapping
	err = json.Unmarshal(body, &gestureMappings)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(resW, "Invalid JSON", http.StatusBadRequest)
		return
	}

	tx, err := models.StartTransaction()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		http.Error(resW, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			err := models.RollbackTransaction(tx)
			if err != nil {
				panic(err)
			}
		}
	}()

	err = models.ClearGestureActionMappings(tx)
	if err != nil {
		log.Printf("Error clearing gesture action mappings: %v", err)
		http.Error(resW, "Error clearing gesture action mappings", http.StatusInternalServerError)
		return
	}

	for _, mapping := range gestureMappings {
		err = models.InsertGestureActionMapping(tx, &mapping)
		if err != nil {
			log.Printf("Error updating gesture mapping: %v", err)
			http.Error(resW, "Error updating gesture mapping", http.StatusInternalServerError)
			return
		}
	}

	err = models.CommitTransaction(tx)
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		http.Error(resW, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	resW.WriteHeader(http.StatusOK)
}

func GetGestureActionMappings(resW http.ResponseWriter, req *http.Request) {
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

	mappings, err := models.GetGestureActionMappings()
	if err != nil {
		log.Printf("Error retrieving gesture action mappings: %v", err)
		http.Error(resW, "Error retrieving gesture action mappings", http.StatusInternalServerError)
		return
	}

	resW.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(resW).Encode(mappings)
	if err != nil {
		http.Error(resW, err.Error(), http.StatusInternalServerError)
	}
}

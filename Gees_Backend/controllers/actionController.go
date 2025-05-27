package controllers

import (
	"Gees_Backend/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

// EditAction an endpoint to edit an action
func EditAction(resW http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		log.Println("Request body is nil")
		http.Error(resW, "Invalid data", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
		http.Error(resW, "Invalid data", http.StatusBadRequest)
		return
	}

	var action models.Action
	err = json.Unmarshal(body, &action)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		http.Error(resW, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if action.ID == nil {
		log.Println("Action ID is nil")
		http.Error(resW, "Invalid action ID", http.StatusBadRequest)
		return
	}

	if action.Name == "" {
		log.Println("Action name is empty")
		http.Error(resW, "Invalid action name", http.StatusBadRequest)
		return
	}

	tx, err := models.StartTransaction()
	if err != nil {
		log.Printf("Error starting transaction: %s", err)
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

	err = models.UpdateAction(tx, &action)
	if err != nil {
		log.Printf("Error updating action: %s", err)
		http.Error(resW, "Error updating action", http.StatusInternalServerError)
		return
	}

	err = models.CommitTransaction(tx)
	if err != nil {
		log.Printf("Error committing transaction: %s", err)
		http.Error(resW, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	resW.WriteHeader(http.StatusOK)
}

// CreateAction an endpoint to create an action returns the action
func CreateAction(resW http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(resW, "Invalid data", http.StatusBadRequest)
		return
	}

	var action models.Action
	err = json.Unmarshal(body, &action)
	if err != nil {
		http.Error(resW, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if action.Name == "" {
		log.Println("Action name is empty")
		http.Error(resW, "Invalid action name", http.StatusBadRequest)
		return
	}

	if action.ID != nil {
		log.Println("Action ID should be nil for new action")
		http.Error(resW, "Action ID should be nil for new action", http.StatusBadRequest)
		return
	}

	tx, err := models.StartTransaction()
	if err != nil {
		log.Printf("Error starting transaction: %s", err)
		http.Error(resW, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			log.Printf("Error rolling back transaction: %v", err)
			err := models.RollbackTransaction(tx)
			if err != nil {
				panic(err)
			}
		}
	}()

	err = models.InsertAction(tx, &action)
	if err != nil {
		log.Printf("Error inserting action: %s", err)
		http.Error(resW, "Error storing data", http.StatusInternalServerError)
		return
	}

	err = models.CommitTransaction(tx)
	if err != nil {
		log.Printf("Error committing transaction: %s", err)
		http.Error(resW, "Error committing transaction", http.StatusInternalServerError)
		return
	}

	resW.WriteHeader(http.StatusOK)
	resW.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(resW).Encode(action)
}

// GetActionByID an endpoint to get an action by ID
func GetActionByID(resW http.ResponseWriter, req *http.Request) {
	actionIDString := req.URL.Path
	actionIDString = actionIDString[len("/action/"):]
	if actionIDString == "" {
		log.Println("Missing action ID")
		http.Error(resW, "Missing action ID", http.StatusBadRequest)
		return
	}
	actionID, err := strconv.Atoi(actionIDString)
	if err != nil {
		log.Printf("Error converting action ID to int: %v", err)
		http.Error(resW, "Invalid action ID", http.StatusBadRequest)
		return
	}
	action, err := models.GetActionByID(actionID)
	if err != nil {
		log.Printf("Error retrieving action: %v", err)
		http.Error(resW, "Error retrieving action", http.StatusInternalServerError)
		return
	}

	resW.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(resW).Encode(action)
	if err != nil {
		log.Printf("Error encoding action to JSON: %v", err)
		http.Error(resW, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetActions an endpoint to get all actions
func GetActions(resW http.ResponseWriter, req *http.Request) {
	tx, err := models.StartTransaction()
	if err != nil {
		http.Error(resW, "Error starting transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			if rollbackErr := models.RollbackTransaction(tx); rollbackErr != nil {
				panic(rollbackErr)
			}
		}
	}()

	actions, err := models.GetAllActions()
	if err != nil {
		http.Error(resW, "Error retrieving actions", http.StatusInternalServerError)
		return
	}

	resW.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(resW).Encode(actions)
	if err != nil {
		http.Error(resW, err.Error(), http.StatusInternalServerError)
	}
}

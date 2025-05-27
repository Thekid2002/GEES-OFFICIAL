package main

import (
	"Gees_Backend/controllers"
	"Gees_Backend/env"
	"Gees_Backend/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resW http.ResponseWriter, req *http.Request) {
		allowedOrigins := []string{"http://localhost:4200", "https://localhost:4200", "http://localhost:5173", "http://my_vue_frontend:80", "http://my_vue_frontend:5173", "http://0.0.0.0:5173",
			"http://0.0.0.0:80", "http://localhost:80", "http://localhost", "http://127.0.0.1:80", "http://127.0.0.1"}
		origin := req.Header.Get("Origin")

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				resW.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		resW.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		resW.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if req.Method == http.MethodOptions {
			resW.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(resW, req)
	})
}

func initialize() {
	env.LoadEnv(nil)
	err := models.Connect()
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = models.InitDatabases()
	if err != nil {
		panic("Failed to initialize databases " + err.Error())
	}

	http.HandleFunc("POST /feature-data", controllers.ArduinoPostFeatureData)
	http.HandleFunc("POST /action", controllers.CreateAction)
	http.HandleFunc("PUT /action/{id}", controllers.EditAction)
	http.HandleFunc("GET /action/{id}", controllers.GetActionByID)
	http.HandleFunc("POST /gesture", controllers.CreateGesture)
	http.HandleFunc("PUT /gesture/{id}", controllers.EditGesture)
	http.HandleFunc("GET /gesture/{id}", controllers.GetGestureByID)
	http.HandleFunc("GET /actions", controllers.GetActions)
	http.HandleFunc("GET /gestures", controllers.GetGestures)
	http.HandleFunc("GET /ws-record-gestures", controllers.RecordGesture)
	http.HandleFunc("GET /ws-validate-gestures", controllers.ValidateGesture)
	http.HandleFunc("PUT /update-gesture-action-mappings", controllers.UpdateGestureActionMappings)
	http.HandleFunc("GET /gesture-action-mappings", controllers.GetGestureActionMappings)

}

func main() {
	initialize()
	hostname := "0.0.0.0"
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:4200", hostname),
		Handler:      corsMiddleware(http.DefaultServeMux), // Apply CORS globally
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	fmt.Printf("Server is listening on http://%s:4200\n", hostname)
	log.Fatal(server.ListenAndServe())
}
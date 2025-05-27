package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func ValidateGesture(resW http.ResponseWriter, req *http.Request) {
	ws, err := upgrader.Upgrade(resW, req, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		http.Error(resW, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}
	handleGestureValidatingWebsocket(ws)
}

func handleGestureValidatingWebsocket(ws *websocket.Conn) {
	GlobalGestureWS = ws
	GlobalReadyForValidatingData = true
}

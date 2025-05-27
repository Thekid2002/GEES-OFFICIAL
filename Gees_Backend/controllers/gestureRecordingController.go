package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/websocket"
)

type MessageTypes string

const (
	COUNTDOWN  MessageTypes = "countdown"
	START      MessageTypes = "start"
	STOP       MessageTypes = "stop"
	ERROR      MessageTypes = "error"
	PREDICTION MessageTypes = "prediction"
)

type WSMessage struct {
	MsgType MessageTypes `json:"msgType"`
	Data    any          `json:"data"`
}

var Msg WSMessage

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections for now, customize as needed
	},
}

func RecordGesture(resW http.ResponseWriter, req *http.Request) {
	ws, err := upgrader.Upgrade(resW, req, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		http.Error(resW, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}
	handleGestureRecordingWebsocket(ws)
}

func handleGestureRecordingWebsocket(ws *websocket.Conn) {
	GlobalGestureWS = ws

	defer func() {
		if err := ws.Close(); err != nil {
			log.Printf("Error closing WebSocket: %v", err)
		}
		GlobalGestureWS = nil
	}()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		log.Printf("Received message: %s", message)

		var id struct {
			ID int `json:"id"`
		}

		if err := json.Unmarshal(message, &id); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// Set the global active gesture ID
		GlobalActiveGestureID = id.ID
		log.Printf("Set active gesture ID to: %d", GlobalActiveGestureID)

		var msg WSMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		done := make(chan bool)
		if GlobalGestureWS != nil && ArduinoConnected() {
			go startCountdown(GlobalGestureWS, 2, done)
			<-done
		}

		if !ArduinoConnected() {
			Msg = WSMessage{
				MsgType: ERROR,
				Data:    "Arduino not connected",
			}
		} else {
			Msg = WSMessage{
				MsgType: START,
				Data:    "go",
			}
		}

		response, err := json.Marshal(Msg)
		if err != nil {
			log.Fatalf("Error marshaling start message: %v", err)
		}
		_ = ws.WriteMessage(websocket.TextMessage, response)
	}
}

func ArduinoConnected() bool {
	return true // Placeholder for actual Arduino connection check
}

func startCountdown(ws *websocket.Conn, duration int, done chan<- bool) {
	for i := duration; i > 0; i-- {
		msg := WSMessage{
			MsgType: COUNTDOWN,
			Data:    strconv.Itoa(i),
		}
		response, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error marshaling countdown message: %v", err)
			return
		}
		if err := ws.WriteMessage(websocket.TextMessage, response); err != nil {
			log.Printf("Error sending countdown message: %v", err)
			return
		}
		time.Sleep(1 * time.Second)
	}

	GlobalReadyForRecordingData = true

	done <- true
}

func WebsocketStopCountdown(ws *websocket.Conn, msg *WSMessage) error {
	if ws == nil {
		log.Println("WebSocket connection is nil")
		return fmt.Errorf("WebSocket connection is nil")
	}
	var responseMsg WSMessage
	if msg != nil {
		responseMsg = *msg
	} else {
		responseMsg = WSMessage{
			MsgType: STOP,
			Data:    "halt",
		}
	}

	response, err := json.Marshal(responseMsg)
	if err != nil {
		log.Fatalf("Error marshaling start message: %v", err)
	}
	_ = ws.WriteMessage(websocket.TextMessage, response)
	GlobalReadyForRecordingData = false
	return nil
}

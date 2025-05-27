#include "WiFiSender.h"
#include "RGB.h"

// Constructor
WiFiSender::WiFiSender(const char* ssid, const char* pass, const char* server, int port) {
  _ssid = ssid;
  _pass = pass;
  _server = server;
  _port = port;
  _lastPostTime = 0;
}

/**
 * Connect to Wi-Fi and server.
 */
void WiFiSender::connect() {
  connectToWiFi();
  connectToServer();
}

// WiFi Connection
void WiFiSender::connectToWiFi() {
  Serial.print("Connecting to Wi-Fi...");
  
  while (WiFi.begin(_ssid, _pass) != WL_CONNECTED) {
    Serial.print(".");
    blinkLED(50, 50, 0, 10);
    delay(1000);
  }
  Serial.println("\nConnected to Wi-Fi!");
}

// Server Connection
void WiFiSender::connectToServer() {
  Serial.print("Connecting to ");
  Serial.println(_server);
  if (_client.connect(_server, _port)) {
    Serial.println("Connected to server!");
  } else {
    Serial.println("Connection failed.");
  }
}

/**
 * Send POST request with JSON data.
 * @param path The path to send the data to.
 * @param jsonData The JSON data to send.
 */
void WiFiSender::post(String path, String& jsonData) {
  if (_client.connected()) {
    Serial.println("Sending data...");
    enableLED(0,0,10);
    //Serial.println(jsonData);

    _client.println("POST " + path + " HTTP/1.1");
    _client.print("Host: ");
    _client.println(_server);
    _client.println("Content-Type: application/json");
    _client.print("Content-Length: ");
    _client.println(jsonData.length());
    _client.println();  // This empty line must be here!

    _client.print(jsonData);

    delay(1000);      // Give time for data to be processed
    _client.flush();  // Ensure buffer is fully sent

    waitForResponse();
    blinkLED(0,10,0,100);
  } else {
    Serial.println("Not connected, retrying...");
    blinkLED(10,0,0,10);
    connectToServer();
  }
}

/**
 * Wait for server response.
 */
String WiFiSender::waitForResponse() {
  unsigned long startTime = millis();
  String response = "";

  while (!_client.available()) {
    if (millis() - startTime > 5000) {
      response = "Server not responding, timed out.";
      return response;
    }
  }

  while (_client.available()) {
    char c = _client.read();
    response += c;
  }

  return response;
}

/**
 * Main loop function to handle data sending and reconnections.
 * @param jsonData The JSON data to send.
 */
void WiFiSender::sendData(String path, String& jsonData) {
  if (!_client.connected()) {
    Serial.println("Connection lost. Reconnecting...");
    _client.stop();
    connectToServer();
  }

  if (millis() - _lastPostTime >= _postInterval) {
    _lastPostTime = millis();
    post(path, jsonData);
  }
}
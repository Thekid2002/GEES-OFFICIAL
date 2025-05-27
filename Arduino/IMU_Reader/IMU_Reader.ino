#include "RGB.h"
#include <Wire.h>
#include <ICM20948_WE.h>
#include <WiFiNINA.h>
#include "WiFiSender.h"
#include "Data.h"
#include <Arduino_LSM6DS3.h>
#include "ReadSensorData.h"
#include "SensorData.h"
#include "Globals.h"
#include "Threshold.h"


const char* ssid = "ArduinoLand";
const char* password = "barnetskriger";
const char* host = "192.168.101.57";
const uint16_t port = 4200;

WiFiSender wifiSender(ssid, password, host, port);

void setup() {
  Serial.begin(9600);
  setupLED();
  Wire.begin();
  wifiSender.connect();

  Serial.println("Position IMU1 flat and don't move it - calibrating...");
  delay(1000);
  if (!IMU.begin()) {
    Serial.println("Failed to initialize IMU!");
  }
  delay(10);
  Serial.println("IMU Calibration Done");
}

void loop() {

  SensorData* sensorData = new SensorData(MAX_DATAPOINTS);

  unsigned long t = millis();

  Serial.println("Waiting for activation threshold...");
  awaitThreshold();

  // Collects reads of sensor data, applies filters, and creates an array of data points
  collectSensorData(sensorData, MAX_DATAPOINTS);
  Serial.println("Time: " + String(millis() - t));

  // Uses the array of data points to create an object that contains:
  // - Features based on the data,
  // - The data points themselves
  // - The size
  Data dataObject = createDataObject(sensorData, MAX_DATAPOINTS);
  String jsonData = dataObject.featuresToJson();

  Serial.println("Attempting to send package");
  wifiSender.sendData("/feature-data", jsonData);
  Serial.println(jsonData);
  delete sensorData;
}
#include "Threshold.h"
#include "RGB.h"
#include "Globals.h"
#include <Arduino_LSM6DS3.h>
#include <Arduino.h>

const unsigned long debounceDelay = 500;

static unsigned long lastDetect = 0;
float ax, ay, az;

bool isThresholdMet(float thresholdForActivation) {

    IMU.readAcceleration(ax, ay, az);

    float magnitude = sqrt(ax*ax + ay*ay + az*az);
  
    unsigned long now = millis();
  
    if (magnitude > thresholdForActivation && (now - lastDetect) > debounceDelay) {
      Serial.println("Motion detected (software threshold)!");
      lastDetect = now;
      return true;
    }
    return false;
}

void awaitThreshold() {
    bool thresholdMet = false;
    while (thresholdMet == false) {
        thresholdMet = isThresholdMet(THRESHOLD_ACTIVATION_VALUE);
        enableLED(0, 10, 5); // Cyan
    }
    disableLED();
}
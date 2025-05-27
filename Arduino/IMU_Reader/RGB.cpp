#include "RGB.h"
#include <WiFiNINA.h>

const int RED_PIN = 26;
const int GREEN_PIN = 25;
const int BLUE_PIN = 27;

void setupLED() {
  pinMode(RED_PIN, OUTPUT);
  pinMode(GREEN_PIN, OUTPUT);
  pinMode(BLUE_PIN, OUTPUT);
}

void enableLED(int red, int green, int blue) {
  WiFiDrv::analogWrite(RED_PIN, red);
  WiFiDrv::analogWrite(GREEN_PIN, green);
  WiFiDrv::analogWrite(BLUE_PIN, blue);
}

void disableLED() {
  WiFiDrv::analogWrite(RED_PIN, 0);
  WiFiDrv::analogWrite(GREEN_PIN, 0);
  WiFiDrv::analogWrite(BLUE_PIN, 0);
}

void blinkLED(int red, int green, int blue, int activeTime) {
  enableLED(red, green, blue);
  delay(activeTime);
  disableLED();
}

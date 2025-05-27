#include "Arduino.h"
#include "ICM20948_WE.h"
#include "KalmanFilter.h"
#include "RGB.h"
#include "SensorReader.h"
#include "SimpleKalmanFilter.h"
#include "WiFiNINA.h"

void delay(unsigned long) {}

unsigned long millis() { return 0; }

long map(long x, long in_min, long in_max, long out_min, long out_max) {
    return (x - in_min) * (out_max - out_min) / (in_max - in_min) + out_min;
}

void setupLED() {}
void enableLED(int r, int g, int b) {}
void disableLED() {}
void blinkLED(int r, int g, int b, int activeTime) {}


void readSensorData(SensorData* s) {
    // Simulate reading sensor data by generating random values
    s->acc.x = static_cast<float>(rand()) / RAND_MAX * 1.0f - 0.5f; 
    s->acc.y = static_cast<float>(rand()) / RAND_MAX * 1.0f - 0.5f;
    s->acc.z = 0.9f + static_cast<float>(rand()) / RAND_MAX * 0.2f; // gravity added
    
    s->gyr.x = static_cast<float>(rand()) / RAND_MAX * 50.0f - 25.0f;  
    s->gyr.y = static_cast<float>(rand()) / RAND_MAX * 50.0f - 25.0f;  
    s->gyr.z = static_cast<float>(rand()) / RAND_MAX * 50.0f - 25.0f;  
}
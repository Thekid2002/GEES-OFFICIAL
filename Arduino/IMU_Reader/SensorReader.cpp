#include "SensorReader.h"
#include <Arduino.h> 
#include <Arduino_LSM6DS3.h>

void readSensorData(SensorData* s) {
    while (!IMU.accelerationAvailable()) {
        delay(1);
    }

    IMU.readAcceleration(s->acc.x, s->acc.y, s->acc.z);
    delay(1);

    IMU.readGyroscope(s->gyr.x, s->gyr.y, s->gyr.z);
    delay(1);
}

#include "ReadSensorData.h"
#include "WiFiSender.h"
#include "Data.h"
#include "FilterArray.h"
#include "FeatureExtraction.h"
#include <SimpleKalmanFilter.h>
#include "KalmanFilter.h"
#include "RGB.h"
#include "Globals.h"
#include "SensorReader.h"
#include "Arduino.h"

SimpleKalmanFilter simpleKalmanFilterAccX(0.001, 0.003, 0.03);
SimpleKalmanFilter simpleKalmanFilterAccZ(0.001, 0.003, 0.03);
SimpleKalmanFilter simpleKalmanFilterAccY(0.001, 0.003, 0.03);
SimpleKalmanFilter simpleKalmanFilterGyrX(0.001, 0.003, 0.03);
SimpleKalmanFilter simpleKalmanFilterGyrY(0.001, 0.003, 0.03);
SimpleKalmanFilter simpleKalmanFilterGyrZ(0.001, 0.003, 0.03);

KalmanFilter* kalmanX;
KalmanFilter* kalmanY;

const int delayBetweenReadings = (GESTURE_READ_TIME - READ_DELAY_OFFSET) / MAX_DATAPOINTS; 

void collectSensorData(SensorData* sensorData, int size) {
    
    kalmanX = new KalmanFilter(0.001, 0.003, 0.03);
    kalmanY = new KalmanFilter(0.001, 0.003, 0.03);

    enableLED(0, 255, 0); // Green
    for (int rawDataCount = 0; rawDataCount < size; rawDataCount++) {
        readSensorData(sensorData);
        delay(delayBetweenReadings); // Amount of datapoints is hardcapped by memory. This spaces out the points in the timeframe.
        applyFilter(sensorData, rawDataCount);
    }

    disableLED();       

    isolateAxes(sensorData, size);
    sensorData->createDataPointArray(POINTS_TO_SEND_COUNT);

    delete kalmanX;
    delete kalmanY;
}

void isolateAxes(SensorData* s, int size) {
    // Isolate accelerometer data
    filterXYZForX(s->accXYZ, size, s->xAxisAcc);
    filterXYZForY(s->accXYZ, size, s->yAxisAcc);
    filterXYZForZ(s->accXYZ, size, s->zAxisAcc);

    // Isolate gyroscope data
    filterXYZForX(s->gyrXYZ, size, s->xAxisGyr);
    filterXYZForY(s->gyrXYZ, size, s->yAxisGyr);
    filterXYZForZ(s->gyrXYZ, size, s->zAxisGyr);
}

Data createDataObject(SensorData* s, int size) {
    // Create Data object with all required parameters
    Data data = Data(
        // Acc Features
        mean(s->xAxisAcc, size), mean(s->yAxisAcc, size), mean(s->zAxisAcc, size),
        variance(s->xAxisAcc, size), variance(s->yAxisAcc, size), variance(s->zAxisAcc, size),
        median(s->xAxisAcc, size), median(s->yAxisAcc, size), median(s->zAxisAcc, size),
        standardDeviation(s->xAxisAcc, size), standardDeviation(s->yAxisAcc, size), standardDeviation(s->zAxisAcc, size),
        skew(s->xAxisAcc, size), skew(s->yAxisAcc, size), skew(s->zAxisAcc, size),
        findMax(s->xAxisAcc, size), findMax(s->yAxisAcc, size), findMax(s->zAxisAcc, size),
        findMin(s->xAxisAcc, size), findMin(s->yAxisAcc, size), findMin(s->zAxisAcc, size),
        // Gyr Features
        mean(s->xAxisGyr, size), mean(s->yAxisGyr, size), mean(s->zAxisGyr, size),
        variance(s->xAxisGyr, size), variance(s->yAxisGyr, size), variance(s->zAxisGyr, size),
        median(s->xAxisGyr, size), median(s->yAxisGyr, size), median(s->zAxisGyr, size),
        standardDeviation(s->xAxisGyr, size), standardDeviation(s->yAxisGyr, size), standardDeviation(s->zAxisGyr, size),
        skew(s->xAxisGyr, size), skew(s->yAxisGyr, size), skew(s->zAxisGyr, size),
        findMax(s->xAxisGyr, size), findMax(s->yAxisGyr, size), findMax(s->zAxisGyr, size),
        findMin(s->xAxisGyr, size), findMin(s->yAxisGyr, size), findMin(s->zAxisGyr, size),
        // Roll and Pitch Features
        mean(s->roll, size), variance(s->roll, size), median(s->roll, size),
        standardDeviation(s->roll, size), skew(s->roll, size), findMax(s->roll, size), findMin(s->roll, size),
        mean(s->pitch, size), variance(s->pitch, size), median(s->pitch, size),
        standardDeviation(s->pitch, size), skew(s->pitch, size), findMax(s->pitch, size), findMin(s->pitch, size),
        s->dataPoints, size);
  return data;
}

/**
 * Apply Kalman filter to the sensor data.
 * @param s Pointer to the SensorData structure.
 * @param currentRawDataCount The current index of the raw data count.
 */
void applyFilter(SensorData* s, int currentRawDataCount) {

    float accPitch = pitch(s->acc.x, s->acc.y, s->acc.z);
    float accRoll  = roll(s->acc.y, s->acc.z);

    float kalPitch = kalmanY->update(accPitch, s->gyr.y);
    float kalRoll = kalmanX->update(accRoll, s->gyr.x);

    s->roll[currentRawDataCount] = kalRoll;
    s->pitch[currentRawDataCount] = kalPitch;

    float estimatedAccX = simpleKalmanFilterAccX.updateEstimate(s->acc.x);
    float estimatedAccY = simpleKalmanFilterAccY.updateEstimate(s->acc.y);
    float estimatedAccZ = simpleKalmanFilterAccZ.updateEstimate(s->acc.z);

    float estimatedGyrX = simpleKalmanFilterGyrX.updateEstimate(s->gyr.x);
    float estimatedGyrY = simpleKalmanFilterGyrY.updateEstimate(s->gyr.y);
    float estimatedGyrZ = simpleKalmanFilterGyrZ.updateEstimate(s->gyr.z);

    s->accXYZ[currentRawDataCount] = { estimatedAccX, estimatedAccY, estimatedAccZ };
    s->gyrXYZ[currentRawDataCount] = { estimatedGyrX, estimatedGyrY, estimatedGyrZ };
}


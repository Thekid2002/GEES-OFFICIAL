#ifndef SENSOR_DATA_H
#define SENSOR_DATA_H

#include "Data.h"
#include "Globals.h"
#include <ICM20948_WE.h>
#include <Arduino.h> // For String class (in tests)

class SensorData {
    public:
        DataPoint* dataPoints;
        xyzFloat acc;
        xyzFloat gyr;
        float* roll;
        float* pitch;
        xyzFloat* accXYZ;
        xyzFloat* gyrXYZ;
        float* xAxisAcc;
        float* yAxisAcc;
        float* zAxisAcc;
        float* xAxisGyr;
        float* yAxisGyr;
        float* zAxisGyr;

        SensorData(int maxDataPoints) {
            this->acc = xyzFloat();
            this->gyr = xyzFloat();
            this->accXYZ = new xyzFloat[maxDataPoints];
            this->gyrXYZ = new xyzFloat[maxDataPoints];
            this->roll = new float[maxDataPoints];
            this->pitch = new float[maxDataPoints];
            this->xAxisAcc = new float[maxDataPoints];
            this->yAxisAcc = new float[maxDataPoints];
            this->zAxisAcc = new float[maxDataPoints];
            this->xAxisGyr = new float[maxDataPoints];
            this->yAxisGyr = new float[maxDataPoints];
            this->zAxisGyr = new float[maxDataPoints];
            this->dataPoints = new DataPoint[POINTS_TO_SEND_COUNT];
        }

        /*
         * Loops through the full array of raw points (accXYZ, gyrXYZ)
         * and creates the datapoint array by taking points evenly the raw data (this is what map() does). 
         */
        void createDataPointArray(int maxDataPoints) {
            int i;
             for (int j = 0; j < POINTS_TO_SEND_COUNT; j++) {
              i = map(j, 0, POINTS_TO_SEND_COUNT - 1, 0, maxDataPoints - 1);  // Map j to the range [0, maxDataPoints-1]
              this->dataPoints[j] = DataPoint(this->accXYZ[i].x, this->accXYZ[i].y, this->accXYZ[i].z, this->gyrXYZ[i].x, this->gyrXYZ[i].y, this->gyrXYZ[i].z);
            }
        }

        ~SensorData() {
            delete[] accXYZ;
            delete[] gyrXYZ;
            delete[] roll;
            delete[] pitch;
            delete[] xAxisAcc;
            delete[] yAxisAcc;
            delete[] zAxisAcc;
            delete[] xAxisGyr;
            delete[] yAxisGyr;
            delete[] zAxisGyr;
            delete[] dataPoints;
        }
};

#endif
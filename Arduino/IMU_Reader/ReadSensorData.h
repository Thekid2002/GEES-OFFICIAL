  
#ifndef READSENSORDATA_H
#define READSENSORDATA_H
#include "WiFiSender.h"
#include "SensorData.h"

void collectSensorData(SensorData* sensorData, int size);
void isolateAxes(SensorData* s, int size);
Data createDataObject(SensorData* s, int size);
void readSensorData(SensorData* sensorData);
void applyFilter(SensorData* s, int currentRawDataCount);

#endif 
#include "ReadSensorData_Test.h"
#include <gtest/gtest.h>
#include "IMU_Reader/SensorData.h"
#include "IMU_Reader/ReadSensorData.h"
#include <iostream>

TEST(ReadSensorDataTest, TestCollectSensorData) {

    const int size = 10; // Example size
    SensorData* sensorData = new SensorData(size);

    collectSensorData(sensorData, size);
    
    ASSERT_NE(sensorData->dataPoints, nullptr);

    ASSERT_EQ(sensorData->xAxisAcc[0], sensorData->accXYZ[0].x);
    std::cout << "xAxisAcc[0]: " << sensorData->xAxisAcc[0] << std::endl;
    ASSERT_EQ(sensorData->yAxisAcc[0], sensorData->accXYZ[0].y);
    std::cout << "yAxisAcc[0]: " << sensorData->yAxisAcc[0] << std::endl;
    ASSERT_EQ(sensorData->zAxisAcc[0], sensorData->accXYZ[0].z);
    std::cout << "zAxisAcc[0]: " << sensorData->zAxisAcc[0] << std::endl;

    ASSERT_EQ(sensorData->xAxisGyr[0], sensorData->gyrXYZ[0].x);
    std::cout << "xAxisGyr[0]: " << sensorData->xAxisGyr[0] << std::endl;
    ASSERT_EQ(sensorData->yAxisGyr[0], sensorData->gyrXYZ[0].y);
    std::cout << "yAxisGyr[0]: " << sensorData->yAxisGyr[0] << std::endl;
    ASSERT_EQ(sensorData->zAxisGyr[0], sensorData->gyrXYZ[0].z);
    std::cout << "zAxisGyr[0]: " << sensorData->zAxisGyr[0] << std::endl;

    ASSERT_NE(sensorData->dataPoints[0].AccX, 0); 
    ASSERT_NE(sensorData->dataPoints[0].AccY, 0);
    ASSERT_NE(sensorData->dataPoints[0].AccZ, 0);
    
    ASSERT_NE(sensorData->roll[0], 0);
    ASSERT_NE(sensorData->pitch[0], 0);
}



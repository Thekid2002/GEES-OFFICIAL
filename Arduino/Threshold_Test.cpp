#include "ReadSensorData_Test.h"
#include <gtest/gtest.h>
#include "IMU_Reader/SensorData.h"
#include "IMU_Reader/ReadSensorData.h"
#include <iostream>
#include "IMU_Reader/Threshold.h"


TEST(ThresholdTest, TestThreshold) {
    const float threshold = 0;  // If the threshold is 0, this should always be true.
    bool result = isThresholdMet(threshold);
    ASSERT_TRUE(result); 
}
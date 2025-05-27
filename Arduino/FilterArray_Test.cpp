//
// Created by Thomas Andersen on 01/05/2025.
//

#include "FilterArray_Test.h"
#include <gtest/gtest.h>
#include <stdlib.h>
#include <ICM20948_WE.h>
#include "IMU_Reader/FilterArray.h"

TEST(FilterArrayTest, FilterXYZForX) {
    xyzFloat values[] = {{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}};
    float xAxis[3];
    filterXYZForX(values, 3, xAxis);
    EXPECT_FLOAT_EQ(xAxis[0], 1.0);
    EXPECT_FLOAT_EQ(xAxis[1], 4.0);
    EXPECT_FLOAT_EQ(xAxis[2], 7.0);
}

TEST(FilterArrayTest, FilterXYZForY) {
    xyzFloat values[] = {{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}};
    float yAxis[3];
    filterXYZForY(values, 3, yAxis);
    EXPECT_FLOAT_EQ(yAxis[0], 2.0);
    EXPECT_FLOAT_EQ(yAxis[1], 5.0);
    EXPECT_FLOAT_EQ(yAxis[2], 8.0);
}

TEST(FilterArrayTest, FilterXYZForZ) {
    xyzFloat values[] = {{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}};
    float zAxis[3];
    filterXYZForZ(values, 3, zAxis);
    EXPECT_FLOAT_EQ(zAxis[0], 3.0);
    EXPECT_FLOAT_EQ(zAxis[1], 6.0);
    EXPECT_FLOAT_EQ(zAxis[2], 9.0);
}



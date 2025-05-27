cmake_minimum_required(VERSION 3.18.4)
project(IMU_Reader)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_BUILD_TYPE Debug)

# Enable testing
enable_testing()

# Add GoogleTest includes and libraries
include_directories(/usr/src/gtest/googletest/include)
link_directories(/usr/lib)

# Include directories 
include_directories(
    ${CMAKE_SOURCE_DIR}/Stubs       # Stub headers first
    ${CMAKE_SOURCE_DIR}             # Project source files after
)

# FeatureExtraction tests
add_executable(FeatureExtraction_Tests
    FeatureExtraction_Test.cpp
    IMU_Reader/FeatureExtraction.cpp
)
target_link_libraries(FeatureExtraction_Tests gtest gtest_main pthread)

# FilterArray tests
add_executable(FilterArray_Tests
    FilterArray_Test.cpp
    IMU_Reader/FilterArray.cpp
    Stubs/Stubs.cpp
)
target_link_libraries(FilterArray_Tests gtest gtest_main pthread)

# ReadSensorData tests
add_executable(ReadSensorData_Tests
    ReadSensorData_Test.cpp
    IMU_Reader/ReadSensorData.cpp
    IMU_Reader/FeatureExtraction.cpp
    IMU_Reader/FilterArray.cpp
    IMU_Reader/WiFiSender.cpp
    Stubs/Stubs.cpp
)

target_link_libraries(ReadSensorData_Tests gtest gtest_main pthread)

# Register tests
add_test(NAME FeatureExtractionTests COMMAND FeatureExtraction_Tests)
add_test(NAME FilterArrayTests COMMAND FilterArray_Tests)
add_test(NAME ReadSensorDataTests COMMAND ReadSensorData_Tests)

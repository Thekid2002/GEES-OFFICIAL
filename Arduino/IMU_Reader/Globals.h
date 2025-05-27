#ifndef GLOBALS_H
#define GLOBALS_H

// I2C addresses
#define IMU1_ADDR 0x68
#define IMU2_ADDR 0x69

// Timing and buffer sizes
#define READ_DELAY_OFFSET 50 // May be unnecessary.
#define MAX_DATAPOINTS 20 // Number of data points, which are included in the feature extraction.
#define GESTURE_READ_TIME 2000 // Time to wait for gesture recognition in milliseconds.
#define POINTS_TO_SEND_COUNT 8 // Number of raw data points to send.
#define THRESHOLD_ACTIVATION_BOOL 1 // Send data based on threshold activation or server request.
#define THRESHOLD_ACTIVATION_VALUE 2.0 // The magnitude threshold for activation. 

#endif 
#pragma once

class KalmanFilter {
    public:
        // Constructor takes process noise, measurement noise, and error estimates
        KalmanFilter(float processNoise, float measurementNoise, float estimatedError) {
            // Ignore these parameters for the mock
        }
    
        // Mock update method that just returns the input value
        float update(float acc, float gyr) {
            // Normally the Kalman filter would use these values to estimate the state
            // For the mock, we'll just return the accelerometer value (acc)
            return acc; // In reality, it would combine acc and gyr for filtering
        }
    };
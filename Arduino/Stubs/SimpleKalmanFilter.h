#pragma once

class SimpleKalmanFilter {
    public:
        SimpleKalmanFilter(float processNoise, float measurementNoise, float estimatedError) {
            // Constructor parameters are ignored in the mock for simplicity
        }
    
        // Mock updateEstimate method
        float updateEstimate(float value) {
            // Just return the input value for testing purposes
            return value;  // Normally, Kalman filter would process the value here
        }
    };
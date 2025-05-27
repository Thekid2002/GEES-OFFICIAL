#ifndef FEATURE_EXTRACTION_H
#define FEATURE_EXTRACTION_H

float mean(const float values[], unsigned char size);
float variance(const float values[], unsigned char size);
float standardDeviation(const float values[], unsigned char size);
float findMax(const float values[], unsigned char size);
float findMin(const float values[], unsigned char size);
float movingAverage(const float values[], unsigned char size);
float median(float values[], unsigned char size);
float skew(float values[], unsigned char size);
float pitch(float x, float y, float z);
float roll(float y, float z);

#endif  // FEATURE_EXTRACTION_H
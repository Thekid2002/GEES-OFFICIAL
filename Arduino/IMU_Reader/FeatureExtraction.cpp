#include "FeatureExtraction.h"
#include <math.h> 
#include <stdlib.h>

/**
 * Calculate the absolute value of a float.
 * @param x The float to calculate the absolute value of.
 * @return The absolute value of the float.
 */
float absie(float x) {
  return (x < 0) ? -x : x;
}

/**
 * Calculate the mean of an array of floats.
 * @param arr The array of floats.
 * @param size The size of the array.
 * @return The mean of the array.
 */
float mean(const float arr[], const unsigned char size) {
  float sum = 0.0;

  for (int i = 0; i < size; i++) {
    sum += arr[i];
  }

  if (size == 0) {
    return 0;
  }

  return sum / size;
}

/**
 * Calculate the variance of an array of floats.
 * @param arr The array of floats.
 * @param size The size of the array.
 * @return The variance of the array.
 */
float variance(const float arr[], const unsigned char size) {
  float m = mean(arr, size);
  float sum = 0.0;

  for (int i = 0; i < size; i++) {
    sum += pow(arr[i] - m, 2);
  }

  if (size == 0) {
    return 0;
  }

  return sum / size;
}

/**
 * Compare function for qsort.
 * @param a The first element to compare.
 * @param b The second element to compare.
 * @return The difference between the two elements.
 */
int compareQSort(const void* a, const void* b) {
  float diff = (*(float*)a - *(float*)b);
  return (diff > 0) - (diff < 0);
}

/**
 * Find the median of an array of integers.
 * @param arr The array of integers.
 * @param n The size of the array.
 */
float median(float values[], const unsigned char size) {
  // This sorts the values array, glhf
  qsort(values, size, sizeof(values[0]), compareQSort);
  return values[size/2];
}

/**
 * Calculate the standard deviation of an array of floats.
 * @param arr The array of floats.
 * @param size The size of the array.
 * @return The standard deviation of the array.
 */
float standardDeviation(const float arr[], const unsigned char size) {
  return sqrt(variance(arr, size));
}


/**
 * Calculate the skewness of an array of floats.
 * @param arr The array of floats.
 * @param size The size of the array.
 * @return The skewness of the array.
 */
float skew(float arr[], const unsigned char size) {
  float std = standardDeviation(arr, size);
  
  if (std <= 0) {
    return 0;
  }

  float mea = mean(arr, size);
  float med = median(arr, size);

  float skew = absie(mea - med);
  skew = skew / std;
  return skew;
}

/**
 * Calculate the max of an array of floats.
 * @param arr The array of floats.
 * @param size The size of the array.
 * @return The max of the array.
 */
float findMax(const float arr[], const unsigned char size) {
  float maxVal = arr[0];
  for (int i = 1; i < size; i++) {
    maxVal = fmax(maxVal, arr[i]);
  }
  return maxVal;
}

/**
 * Calculate the min of an array of floats.
 * @param arr The array of floats.
 * @param size The size of the array.
 * @return The min of the array.
 */
float findMin(const float arr[], const unsigned char size) {
  float minVal = arr[0];
  for (int i = 1; i < size; i++) {
    minVal = fmin(minVal, arr[i]);  // Using fmin
  }
  return minVal;
}

/**
 * Calculate the moving average of an array of floats.
 * @param arr The array of floats.
 * @param size The size of the array.
 * @return The moving average of the array.
 */
float movingAverage(const float arr[], const unsigned char size) {
  float average = 0.0;
  for (int i = 0; i < size; i++) {
    average += arr[i];
  }

  if (size == 0) {
    return 0;
  }

  return average/= size;
}

/**
 * Calculate the pitch of an object based on its accelerometer data.
 * @param x The x-axis accelerometer data.
 * @param y The y-axis accelerometer data.
 * @param z The z-axis accelerometer data.
 * @return The pitch in degrees.
 */
float pitch(float x, float y, float z) {
  return -(atan2(x, sqrt(y * y + z * z)) * 180.0) / M_PI;
}

/**
 * Calculate the roll of an object based on its accelerometer data.
 * @param y The y-axis accelerometer data.
 * @param z The z-axis accelerometer data.
 * @return The roll in degrees.
 */
float roll(float y, float z) {
  return (atan2(y, z) * 180.0) / M_PI;
}
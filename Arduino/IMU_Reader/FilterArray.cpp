#include "FilterArray.h"

/**
 * Filter an array of xyzFloat values to get the x-axis values.
 * @param values The array of xyzFloat values.
 * @param size The size of the array.
 * @param xAxis The array to store the filtered x-axis values.
 */
void filterXYZForX(const xyzFloat values[], const unsigned char size, float xAxis[]) {
  for (int i = 0; i < size; i++) {
    xAxis[i] = values[i].x;
  }
}

/**
 * Filter an array of xyzFloat values to get the y-axis values.
 * @param values The array of xyzFloat values.
 * @param size The size of the array.
 * @param yAxis The array to store the filtered y-axis values.
 */
void filterXYZForY(const xyzFloat values[], const unsigned char size, float yAxis[]) {
  for (int i = 0; i < size; i++) {
    yAxis[i] = values[i].y;
  }
}

/**
 * Filter an array of xyzFloat values to get the z-axis values.
 * @param values The array of xyzFloat values.
 * @param size The size of the array.
 * @param zAxis The array to store the filtered z-axis values.
 */
void filterXYZForZ(const xyzFloat values[], const unsigned char size, float zAxis[]) {
  for (int i = 0; i < size; i++) {
    zAxis[i] = values[i].z;
  }
}
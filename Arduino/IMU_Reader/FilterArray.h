//
// Created by Thomas Andersen on 28/03/2025.
//

#ifndef FILTERARRAY_H
#define FILTERARRAY_H
#include <ICM20948_WE.h>


void filterXYZForX(const xyzFloat values[], unsigned char size, float* xAxis);

void filterXYZForY(const xyzFloat values[], unsigned char size, float* yAxis);

void filterXYZForZ(const xyzFloat values[], unsigned char size, float* zAxis);


#endif //FILTERARRAY_H

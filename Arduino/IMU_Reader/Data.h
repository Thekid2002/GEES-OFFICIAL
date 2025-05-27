  #ifndef DATA_H
  #define DATA_H

  #include "Globals.h"
  #include <Arduino.h> // For String class (in tests)

  class DataPoint {
  public:
    float AccX;
    float AccY;
    float AccZ;
    float GyrX;
    float GyrY;
    float GyrZ;

    void toJson(String& json) {
      json += "{";
      json += "\"AX\":" + String(AccX, 2) + ",";
      json += "\"AY\":" + String(AccY, 2) + ",";
      json += "\"AZ\":" + String(AccZ, 2) + ",";
      json += "\"GX\":" + String(GyrX, 2) + ",";
      json += "\"GY\":" + String(GyrY, 2) + ",";
      json += "\"GZ\":" + String(GyrZ, 2);
      json += "}";
    }

    DataPoint(float accX, float accY, float accZ, float gyrX, float gyrY, float gyrZ) {
      AccX = accX;
      AccY = accY;
      AccZ = accZ;
      GyrX = gyrX;
      GyrY = gyrY;
      GyrZ = gyrZ;
    }

    DataPoint() {
      AccX = 0;
      AccY = 0;
      AccZ = 0;
      GyrX = 0;
      GyrY = 0;
      GyrZ = 0;
    }
  };

  class Data {
  public:
    float AccMeanX;
    float AccMeanY;
    float AccMeanZ;
    float AccVarianceX;
    float AccVarianceY;
    float AccVarianceZ;
    float AccMedianX;
    float AccMedianY;
    float AccMedianZ;
    float AccStdDevX;
    float AccStdDevY;
    float AccStdDevZ;
    float AccSkewX;
    float AccSkewY;
    float AccSkewZ;
    float AccMaxX;
    float AccMaxY;
    float AccMaxZ;
    float AccMinX;
    float AccMinY;
    float AccMinZ;

    float GyrMeanX;
    float GyrMeanY;
    float GyrMeanZ;
    float GyrVarianceX;
    float GyrVarianceY;
    float GyrVarianceZ;
    float GyrMedianX;
    float GyrMedianY;
    float GyrMedianZ;
    float GyrStdDevX;
    float GyrStdDevY;
    float GyrStdDevZ;
    float GyrSkewX;
    float GyrSkewY;
    float GyrSkewZ;
    float GyrMaxX;
    float GyrMaxY;
    float GyrMaxZ;
    float GyrMinX;
    float GyrMinY;
    float GyrMinZ;

    float RollMean;
    float RollVariance;
    float RollMedian;
    float RollStdDev;
    float RollSkew;
    float RollMax;
    float RollMin;

    float PitchMean;
    float PitchVariance;
    float PitchMedian;
    float PitchStdDev;
    float PitchSkew;
    float PitchMax;
    float PitchMin;

    DataPoint* DataPoints;
    void dataPointsToJson(String& json) {
      json += "[";         
      for (int i = 0; i < POINTS_TO_SEND_COUNT; i++) {
        DataPoints[i].toJson(json);
        if (i < POINTS_TO_SEND_COUNT-1) {  // Avoid trailing comma
          json += ",";
        }
      }
      json += "]";
    }

    String featuresToJson() {
      String json;
      json.reserve(2800); // DONT TOUCH (unless it doesnt work)

      json = "{";
      json += "\"AccMeanX\":" + String(AccMeanX, 2) + ",";
      json += "\"AccMeanY\":" + String(AccMeanY, 2) + ",";
      json += "\"AccMeanZ\":" + String(AccMeanZ, 2) + ",";
      json += "\"AccVarX\":" + String(AccVarianceX, 2) + ",";
      json += "\"AccVarY\":" + String(AccVarianceY, 2) + ",";
      json += "\"AccVarZ\":" + String(AccVarianceZ, 2) + ",";
      json += "\"AccMedX\":" + String(AccMedianX, 2) + ",";
      json += "\"AccMedY\":" + String(AccMedianY, 2) + ",";
      json += "\"AccMedZ\":" + String(AccMedianZ, 2) + ",";
      json += "\"AccStdX\":" + String(AccStdDevX, 2) + ",";
      json += "\"AccStdY\":" + String(AccStdDevY, 2) + ",";
      json += "\"AccStdZ\":" + String(AccStdDevZ, 2) + ",";
      json += "\"AccSkewX\":" + String(AccSkewX, 3) + ",";
      json += "\"AccSkewY\":" + String(AccSkewY, 3) + ",";
      json += "\"AccSkewZ\":" + String(AccSkewZ, 3) + ",";
      json += "\"AccMaxX\":" + String(AccMaxX, 2) + ",";
      json += "\"AccMaxY\":" + String(AccMaxY, 2) + ",";
      json += "\"AccMaxZ\":" + String(AccMaxZ, 2) + ",";
      json += "\"AccMinX\":" + String(AccMinX, 2) + ",";
      json += "\"AccMinY\":" + String(AccMinY, 2) + ",";
      json += "\"AccMinZ\":" + String(AccMinZ, 2) + ",";
      json += "\"GyrMeanX\":" + String(GyrMeanX, 2) + ",";
      json += "\"GyrMeanY\":" + String(GyrMeanY, 2) + ",";
      json += "\"GyrMeanZ\":" + String(GyrMeanZ, 2) + ",";
      json += "\"GyrVarX\":" + String(GyrVarianceX, 2) + ",";
      json += "\"GyrVarY\":" + String(GyrVarianceY, 2) + ",";
      json += "\"GyrVarZ\":" + String(GyrVarianceZ, 2) + ",";
      json += "\"GyrMedX\":" + String(GyrMedianX, 2) + ",";
      json += "\"GyrMedY\":" + String(GyrMedianY, 2) + ",";
      json += "\"GyrMedZ\":" + String(GyrMedianZ, 2) + ",";
      json += "\"GyrStdX\":" + String(GyrStdDevX, 2) + ",";
      json += "\"GyrStdY\":" + String(GyrStdDevY, 2) + ",";
      json += "\"GyrStdZ\":" + String(GyrStdDevZ, 2) + ",";
      json += "\"GyrSkewX\":" + String(GyrSkewX, 3) + ",";
      json += "\"GyrSkewY\":" + String(GyrSkewY, 3) + ",";
      json += "\"GyrSkewZ\":" + String(GyrSkewZ, 3) + ",";
      json += "\"GyrMaxX\":" + String(GyrMaxX, 2) + ",";
      json += "\"GyrMaxY\":" + String(GyrMaxY, 2) + ",";
      json += "\"GyrMaxZ\":" + String(GyrMaxZ, 2) + ",";
      json += "\"GyrMinX\":" + String(GyrMinX, 2) + ",";
      json += "\"GyrMinY\":" + String(GyrMinY, 2) + ",";
      json += "\"GyrMinZ\":" + String(GyrMinZ, 2) + ",";
      json += "\"RollMean\":" + String(RollMean, 2) + ",";
      json += "\"RollVar\":" + String(RollVariance, 2) + ",";
      json += "\"RollMed\":" + String(RollMedian, 2) + ",";
      json += "\"RollStd\":" + String(RollStdDev, 2) + ",";
      json += "\"RollSkew\":" + String(RollSkew, 3) + ",";
      json += "\"RollMax\":" + String(RollMax, 2) + ",";
      json += "\"RollMin\":" + String(RollMin, 2) + ",";
      json += "\"PitchMean\":" + String(PitchMean, 2) + ",";
      json += "\"PitchVar\":" + String(PitchVariance, 2) + ",";
      json += "\"PitchMed\":" + String(PitchMedian, 2) + ",";
      json += "\"PitchStd\":" + String(PitchStdDev, 2) + ",";
      json += "\"PitchSkew\":" + String(PitchSkew, 3) + ",";
      json += "\"PitchMax\":" + String(PitchMax, 2) + ",";
      json += "\"PitchMin\":" + String(PitchMin, 2) + ",";
      json += "\"DataPoints\":"; 
      dataPointsToJson(json);
      json += "}";
      return json;
    }

    Data(float accMeanX, float accMeanY, float accMeanZ,
        float accVarianceX, float accVarianceY, float accVarianceZ,
        float accMedianX, float accMedianY, float accMedianZ,
        float accStdDevX, float accStdDevY, float accStdDevZ,
        float accSkewX, float accSkewY, float accSkewZ,
        float accMaxX, float accMaxY, float accMaxZ,
        float accMinX, float accMinY, float accMinZ,
        float gyrMeanX, float gyrMeanY, float gyrMeanZ,
        float gyrVarianceX, float gyrVarianceY, float gyrVarianceZ,
        float gyrMedianX, float gyrMedianY, float gyrMedianZ,
        float gyrStdDevX, float gyrStdDevY, float gyrStdDevZ,
        float gyrSkewX, float gyrSkewY, float gyrSkewZ,
        float gyrMaxX, float gyrMaxY, float gyrMaxZ,
        float gyrMinX, float gyrMinY, float gyrMinZ,
        float rollMean, float rollVariance, float rollMedian,
        float rollStdDev, float rollSkew, float rollMax, float rollMin,
        float pitchMean, float pitchVariance, float pitchMedian,
        float pitchStdDev, float pitchSkew, float pitchMax, float pitchMin,
        DataPoint* dataPoints, int dataPointsCount) {
      AccMeanX = accMeanX;
      AccMeanY = accMeanY;
      AccMeanZ = accMeanZ;
      AccVarianceX = accVarianceX;
      AccVarianceY = accVarianceY;
      AccVarianceZ = accVarianceZ;
      AccMedianX = accMedianX;
      AccMedianY = accMedianY;
      AccMedianZ = accMedianZ;
      AccStdDevX = accStdDevX;
      AccStdDevY = accStdDevY;
      AccStdDevZ = accStdDevZ;
      AccSkewX = accSkewX;
      AccSkewY = accSkewY;
      AccSkewZ = accSkewZ;
      AccMaxX = accMaxX;
      AccMaxY = accMaxY;
      AccMaxZ = accMaxZ;
      AccMinX = accMinX;
      AccMinY = accMinY;
      AccMinZ = accMinZ;

      GyrMeanX = gyrMeanX;
      GyrMeanY = gyrMeanY;
      GyrMeanZ = gyrMeanZ;
      GyrVarianceX = gyrVarianceX;
      GyrVarianceY = gyrVarianceY;
      GyrVarianceZ = gyrVarianceZ;
      GyrMedianX = gyrMedianX;
      GyrMedianY = gyrMedianY;
      GyrMedianZ = gyrMedianZ;
      GyrStdDevX = gyrStdDevX;
      GyrStdDevY = gyrStdDevY;
      GyrStdDevZ = gyrStdDevZ;
      GyrSkewX = gyrSkewX;
      GyrSkewY = gyrSkewY;
      GyrSkewZ = gyrSkewZ;
      GyrMaxX = gyrMaxX;
      GyrMaxY = gyrMaxY;
      GyrMaxZ = gyrMaxZ;
      GyrMinX = gyrMinX;
      GyrMinY = gyrMinY;
      GyrMinZ = gyrMinZ;

      RollMean = rollMean;
      RollVariance = rollVariance;
      RollMedian = rollMedian;
      RollStdDev = rollStdDev;
      RollSkew = rollSkew;
      RollMax = rollMax;
      RollMin = rollMin;

      PitchMean = pitchMean;
      PitchVariance = pitchVariance;
      PitchMedian = pitchMedian;
      PitchStdDev = pitchStdDev;
      PitchSkew = pitchSkew;
      PitchMax = pitchMax;
      PitchMin = pitchMin;
      DataPoints = dataPoints;
    }
  };

  #endif
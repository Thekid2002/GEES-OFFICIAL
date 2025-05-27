#include "IMU_Reader/FeatureExtraction.h"
#include <math.h>
#include <gtest/gtest.h>
#include <cstdlib>
#include <ctime>

float generateRandomFloat() {
    return static_cast<float>(std::rand()) / static_cast<float>(RAND_MAX);
}

void generateRandomArray(float* arr, int size) {
    for (int i = 0; i < size; ++i) {
        arr[i] = generateRandomFloat();
    }
}

TEST(FeatureExtractionFuzzTest, MeanFuzzTest) {
    std::srand(std::time(nullptr));
    const int size = generateRandomFloat();
    float arr[size];
    generateRandomArray(arr, size);

    try {
        mean(arr, size);
    } catch (const std::exception& e) {
        std::cerr << "Exception in mean(): " << e.what() << " | Input size: " << size << std::endl;
        for (int i = 0; i < size; ++i) {
            std::cerr << "arr[" << i << "] = " << arr[i] << std::endl;
        }
        throw;
    }
}

TEST(FeatureExtractionFuzzTest, VarianceFuzzTest) {
    std::srand(std::time(nullptr));
    const int size = generateRandomFloat();
    float arr[size];
    generateRandomArray(arr, size);

    try {
        variance(arr, size);
    } catch (const std::exception& e) {
        std::cerr << "Exception in variance(): " << e.what() << " | Input size: " << size << std::endl;
        for (int i = 0; i < size; ++i) {
            std::cerr << "arr[" << i << "] = " << arr[i] << std::endl;
        }
        throw;
    }
}

TEST(FeatureExtractionFuzzTest, StandardDeviationFuzzTest) {
    std::srand(std::time(nullptr));
    const int size = generateRandomFloat();
    float arr[size];
    generateRandomArray(arr, size);

    try {
        standardDeviation(arr, size);
    } catch (const std::exception& e) {
        std::cerr << "Exception in standardDeviation(): " << e.what() << " | Input size: " << size << std::endl;
        for (int i = 0; i < size; ++i) {
            std::cerr << "arr[" << i << "] = " << arr[i] << std::endl;
        }
        throw;
    }
}

TEST(FeatureExtractionFuzzTest, MedianFuzzTest) {
    std::srand(std::time(nullptr));
    const int size = generateRandomFloat();
    float arr[size];
    generateRandomArray(arr, size);

    try {
        median(arr, size);
    } catch (const std::exception& e) {
        std::cerr << "Exception in median(): " << e.what() << " | Input size: " << size << std::endl;
        for (int i = 0; i < size; ++i) {
            std::cerr << "arr[" << i << "] = " << arr[i] << std::endl;
        }
        throw;
    }
}

TEST(FeatureExtractionFuzzTest, SkewFuzzTest) {
    std::srand(std::time(nullptr));
    const int size = generateRandomFloat();
    float arr[size];
    generateRandomArray(arr, size);

    try {
        skew(arr, size);
    } catch (const std::exception& e) {
        std::cerr << "Exception in skew(): " << e.what() << " | Input size: " << size << std::endl;
        for (int i = 0; i < size; ++i) {
            std::cerr << "arr[" << i << "] = " << arr[i] << std::endl;
        }
        throw;
    }
}

TEST(FeatureExtractionFuzzTest, FindMaxFuzzTest) {
    std::srand(std::time(nullptr));
    const int size = generateRandomFloat();
    float arr[size];
    generateRandomArray(arr, size);

    try {
        findMax(arr, size);
    } catch (const std::exception& e) {
        std::cerr << "Exception in findMax(): " << e.what() << " | Input size: " << size << std::endl;
        for (int i = 0; i < size; ++i) {
            std::cerr << "arr[" << i << "] = " << arr[i] << std::endl;
        }
        throw;
    }
}

TEST(FeatureExtractionFuzzTest, FindMinFuzzTest) {
    std::srand(std::time(nullptr));
    const int size = generateRandomFloat();
    float arr[size];
    generateRandomArray(arr, size);

    try {
        findMin(arr, size);
    } catch (const std::exception& e) {
        std::cerr << "Exception in findMin(): " << e.what() << " | Input size: " << size << std::endl;
        for (int i = 0; i < size; ++i) {
            std::cerr << "arr[" << i << "] = " << arr[i] << std::endl;
        }
        throw;
    }
}

TEST(FeatureExtractionTest, MeanTest) {
    float arr[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    EXPECT_FLOAT_EQ(mean(arr, 5), 3.0);
}

TEST(FeatureExtractionTest, VarianceTest) {
    float arr[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    EXPECT_FLOAT_EQ(variance(arr, 5), 2.0);
}

TEST(FeatureExtractionTest, StandardDeviationTest) {
    float arr[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    EXPECT_FLOAT_EQ(standardDeviation(arr, 5), sqrt(2.0));
}

TEST(FeatureExtractionTest, MedianTestOdd) {
    float arr[] = {5.0, 1.0, 3.0};
    EXPECT_FLOAT_EQ(median(arr, 3), 3.0);
}

TEST(FeatureExtractionTest, MedianTestEven) {
    float arr[] = {5.0, 1.0, 3.0, 2.0};
    EXPECT_FLOAT_EQ(median(arr, 4), 3.0);
}

TEST(FeatureExtractionTest, SkewTest) {
    float arr[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    EXPECT_FLOAT_EQ(skew(arr, 5), 0.0);
}

TEST(FeatureExtractionTest, FindMaxTest) {
    float arr[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    EXPECT_FLOAT_EQ(findMax(arr, 5), 5.0);
}

TEST(FeatureExtractionTest, FindMinTest) {
    float arr[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    EXPECT_FLOAT_EQ(findMin(arr, 5), 1.0);
}

TEST(FeatureExtractionTest, MovingAverageTest) {
    float arr[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    EXPECT_FLOAT_EQ(movingAverage(arr, 5), 3.0);
}
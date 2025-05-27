package models

import (
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func randomFloat() float64 {
	return rand.Float64() * 100
}

func fillStructWithRandomValues(s interface{}) {
	rand.Seed(time.Now().UnixNano())

	v := reflect.ValueOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.CanSet() {
			if field.Kind() == reflect.Float64 {
				field.SetFloat(randomFloat())
			} else if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Float64 {
				ptr := randomFloat()
				field.Set(reflect.ValueOf(&ptr))
			}
		}
	}
}

func CompareDataPoint(retrieved DataPoint, expected DataPoint) bool {
	return retrieved == expected
}

func CompareFeatureData(retrieved FeatureData, expected FeatureData) bool {
	return *retrieved.AccMeanX == *expected.AccMeanX &&
		*retrieved.AccMeanY == *expected.AccMeanY &&
		*retrieved.AccMeanZ == *expected.AccMeanZ &&
		*retrieved.AccVarianceX == *expected.AccVarianceX &&
		*retrieved.AccVarianceY == *expected.AccVarianceY &&
		*retrieved.AccVarianceZ == *expected.AccVarianceZ &&
		*retrieved.AccMedianX == *expected.AccMedianX &&
		*retrieved.AccMedianY == *expected.AccMedianY &&
		*retrieved.AccMedianZ == *expected.AccMedianZ &&
		*retrieved.AccStdDevX == *expected.AccStdDevX &&
		*retrieved.AccStdDevY == *expected.AccStdDevY &&
		*retrieved.AccStdDevZ == *expected.AccStdDevZ &&
		*retrieved.AccSkewX == *expected.AccSkewX &&
		*retrieved.AccSkewY == *expected.AccSkewY &&
		*retrieved.AccSkewZ == *expected.AccSkewZ &&
		*retrieved.AccMaxX == *expected.AccMaxX &&
		*retrieved.AccMaxY == *expected.AccMaxY &&
		*retrieved.AccMaxZ == *expected.AccMaxZ &&
		*retrieved.AccMinX == *expected.AccMinX &&
		*retrieved.AccMinY == *expected.AccMinY &&
		*retrieved.AccMinZ == *expected.AccMinZ &&
		*retrieved.GyrMeanX == *expected.GyrMeanX &&
		*retrieved.GyrMeanY == *expected.GyrMeanY &&
		*retrieved.GyrMeanZ == *expected.GyrMeanZ &&
		*retrieved.GyrVarianceX == *expected.GyrVarianceX &&
		*retrieved.GyrVarianceY == *expected.GyrVarianceY &&
		*retrieved.GyrVarianceZ == *expected.GyrVarianceZ &&
		*retrieved.GyrMedianX == *expected.GyrMedianX &&
		*retrieved.GyrMedianY == *expected.GyrMedianY &&
		*retrieved.GyrMedianZ == *expected.GyrMedianZ &&
		*retrieved.GyrStdDevX == *expected.GyrStdDevX &&
		*retrieved.GyrStdDevY == *expected.GyrStdDevY &&
		*retrieved.GyrStdDevZ == *expected.GyrStdDevZ &&
		*retrieved.GyrSkewX == *expected.GyrSkewX &&
		*retrieved.GyrSkewY == *expected.GyrSkewY &&
		*retrieved.GyrSkewZ == *expected.GyrSkewZ &&
		*retrieved.GyrMaxX == *expected.GyrMaxX &&
		*retrieved.GyrMaxY == *expected.GyrMaxY &&
		*retrieved.GyrMaxZ == *expected.GyrMaxZ &&
		*retrieved.GyrMinX == *expected.GyrMinX &&
		*retrieved.GyrMinY == *expected.GyrMinY &&
		*retrieved.GyrMinZ == *expected.GyrMinZ &&
		*retrieved.RollMean == *expected.RollMean &&
		*retrieved.RollVariance == *expected.RollVariance &&
		*retrieved.RollMedian == *expected.RollMedian &&
		*retrieved.RollStdDev == *expected.RollStdDev &&
		*retrieved.RollSkew == *expected.RollSkew &&
		*retrieved.RollMax == *expected.RollMax &&
		*retrieved.RollMin == *expected.RollMin &&
		*retrieved.PitchMean == *expected.PitchMean &&
		*retrieved.PitchVariance == *expected.PitchVariance &&
		*retrieved.PitchMedian == *expected.PitchMedian &&
		*retrieved.PitchStdDev == *expected.PitchStdDev &&
		*retrieved.PitchSkew == *expected.PitchSkew &&
		*retrieved.PitchMax == *expected.PitchMax &&
		*retrieved.PitchMin == *expected.PitchMin &&
		*retrieved.GestureID == *expected.GestureID
}

func insertTestGesture() Gesture {
	tx, err := StartTransaction()
	if err != nil {
		log.Fatal(err)
	}

	gesture := Gesture{
		Name: "Test Gesture",
	}

	err = InsertGesture(tx, &gesture)
	if err != nil {
		log.Fatal(err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}

	return gesture
}

func TestInsertAndRetrieveFeatureData(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	testGesture := insertTestGesture()

	testFeatureData := &FeatureData{}
	fillStructWithRandomValues(testFeatureData)
	testFeatureData.GestureID = testGesture.ID

	err = InsertFeatureData(tx, testFeatureData, *testGesture.ID)
	if err != nil {
		t.Fatalf("Failed to insert feature data: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	retrievedFeatureData, err := GetFeatureDataByID(testFeatureData.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve feature data: %v", err)
	}

	if !CompareFeatureData(*retrievedFeatureData, *testFeatureData) {
		t.Fatalf("Retrieved feature data does not match expected values.\nGot: %+v\nExpected: %+v", retrievedFeatureData, *testFeatureData)
	}
}

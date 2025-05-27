package models

import (
	"testing"
)

func TestConvertDataPointsToJSON(t *testing.T) {
	dataPoints := []DataPoint{
		{
			AccX: 1.0,
			AccY: 2.0,
			AccZ: 3.0,
			GyrX: 4.0,
			GyrY: 5.0,
			GyrZ: 6.0,
		},
		{
			AccX: 7.0,
			AccY: 8.0,
			AccZ: 9.0,
			GyrX: 10.0,
			GyrY: 11.0,
			GyrZ: 12.0,
		},
	}

	jsonData := ConvertDataPointsToJSON(dataPoints)

	if len(jsonData) == 0 {
		t.Fatal("JSON data is empty")
	}

}

func TestGetDataPoints(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	_, err = GetDataPoints(tx)
	if err != nil {
		t.Fatalf("Failed to get data points: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}

func TestInsertDataPoint(t *testing.T) {
	tx, err := StartTransaction()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}

	dataPoint := DataPoint{
		AccX: 1.0,
		AccY: 2.0,
		AccZ: 3.0,
		GyrX: 4.0,
		GyrY: 5.0,
		GyrZ: 6.0,
	}

	err = InsertDataPoint(tx, &dataPoint)
	if err != nil {
		t.Fatalf("Failed to insert data point: %v", err)
	}

	err = CommitTransaction(tx)
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}

package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

// DataPoint struct
type DataPoint struct {
	ID            *int    `json:"ID"`
	AccX          float64 `json:"AX"`
	AccY          float64 `json:"AY"`
	AccZ          float64 `json:"AZ"`
	GyrX          float64 `json:"GX"`
	GyrY          float64 `json:"GY"`
	GyrZ          float64 `json:"GZ"`
	FeatureDataID *int    `json:"FeatureDataID"`
}

// ConvertDataPointsToJSON converts a slice of DataPoint to a JSON object
func ConvertDataPointsToJSON(dataPoint []DataPoint) string {
	data := make(map[string]interface{})
	data["DataPoints"] = dataPoint

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error converting data points to JSON: %v", err)
		return ""
	}

	return string(jsonData)
}

// CreateDataPointTable creates the data_points table if it does not exist
func CreateDataPointTable(tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS data_points (
	    id SERIAL PRIMARY KEY,
		acc_x FLOAT NOT NULL,
		acc_y FLOAT NOT NULL,
		acc_z FLOAT NOT NULL,
		gyro_x FLOAT NOT NULL,
		gyro_y FLOAT NOT NULL,
		gyro_z FLOAT NOT NULL,
		feature_data_id INT REFERENCES feature_data(id) ON DELETE CASCADE
	)`
	_, err := tx.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("Error creating data_points table: %v\n", err)
	}
	log.Printf("Created DataPoint table successfully")
	return nil
}

func InsertDataPoint(tx pgx.Tx, dataPoint *DataPoint) error {
	var data = *dataPoint
	query := `
	INSERT INTO data_points (acc_x, acc_y, acc_z, gyro_x, gyro_y, gyro_z, feature_data_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := tx.QueryRow(context.Background(), query,
		data.AccX,
		data.AccY,
		data.AccZ,
		data.GyrX,
		data.GyrY,
		data.GyrZ,
		data.FeatureDataID,
	).Scan(&dataPoint.ID)
	if err != nil {
		return fmt.Errorf("error inserting data point: %v", err)
	}
	log.Printf("Inserted DataPoint successfully")
	return nil
}

func GetDataPointByID(tx pgx.Tx, id int) (DataPoint, error) {
	query := `
	SELECT id, acc_x, acc_y, acc_z, gyro_x, gyro_y, gyro_z
	FROM data_points WHERE id = $1`
	var dataPoint DataPoint
	err := tx.QueryRow(context.Background(), query, id).Scan(
		&dataPoint.ID,
		&dataPoint.AccX,
		&dataPoint.AccY,
		&dataPoint.AccZ,
		&dataPoint.GyrX,
		&dataPoint.GyrY,
		&dataPoint.GyrZ,
	)
	if err != nil {
		return dataPoint, fmt.Errorf("error retrieving data point: %v", err)
	}
	return dataPoint, nil
}

func GetDataPointsByFeatureDataID(tx pgx.Tx, featureDataID int) ([]DataPoint, error) {
	query := `
	SELECT id, acc_x, acc_y, acc_z, gyro_x, gyro_y, gyro_z
	FROM data_points WHERE feature_data_id = $1`
	rows, err := tx.Query(context.Background(), query, featureDataID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data points: %v", err)
	}
	defer rows.Close()

	var dataPoints []DataPoint
	for rows.Next() {
		var dataPoint DataPoint
		err := rows.Scan(&dataPoint.ID, &dataPoint.AccX, &dataPoint.AccY, &dataPoint.AccZ,
			&dataPoint.GyrX, &dataPoint.GyrY, &dataPoint.GyrZ)
		if err != nil {
			return nil, fmt.Errorf("error scanning data point: %v", err)
		}
		dataPoints = append(dataPoints, dataPoint)
	}
	return dataPoints, nil
}

func GetDataPoints(tx pgx.Tx) ([]DataPoint, error) {
	query := `
	SELECT id, acc_x, acc_y, acc_z, gyro_x, gyro_y, gyro_z
	FROM data_points`
	rows, err := tx.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data points: %v", err)
	}
	defer rows.Close()

	var dataPoints []DataPoint
	for rows.Next() {
		var dataPoint DataPoint
		err := rows.Scan(&dataPoint.ID, &dataPoint.AccX, &dataPoint.AccY, &dataPoint.AccZ,
			&dataPoint.GyrX, &dataPoint.GyrY, &dataPoint.GyrZ)
		if err != nil {
			return nil, fmt.Errorf("error scanning data point: %v", err)
		}
		dataPoints = append(dataPoints, dataPoint)
	}
	return dataPoints, nil
}

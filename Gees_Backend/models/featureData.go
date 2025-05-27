package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

type FeatureData struct {
	ID            *int        `json:"ID,omitempty"`
	AccMeanX      *float64    `json:"AccMeanX"`
	AccMeanY      *float64    `json:"AccMeanY"`
	AccMeanZ      *float64    `json:"AccMeanZ"`
	AccVarianceX  *float64    `json:"AccVarX"`
	AccVarianceY  *float64    `json:"AccVarY"`
	AccVarianceZ  *float64    `json:"AccVarZ"`
	AccMedianX    *float64    `json:"AccMedX"`
	AccMedianY    *float64    `json:"AccMedY"`
	AccMedianZ    *float64    `json:"AccMedZ"`
	AccStdDevX    *float64    `json:"AccStdX"`
	AccStdDevY    *float64    `json:"AccStdY"`
	AccStdDevZ    *float64    `json:"AccStdZ"`
	AccSkewX      *float64    `json:"AccSkewX"`
	AccSkewY      *float64    `json:"AccSkewY"`
	AccSkewZ      *float64    `json:"AccSkewZ"`
	AccMaxX       *float64    `json:"AccMaxX"`
	AccMaxY       *float64    `json:"AccMaxY"`
	AccMaxZ       *float64    `json:"AccMaxZ"`
	AccMinX       *float64    `json:"AccMinX"`
	AccMinY       *float64    `json:"AccMinY"`
	AccMinZ       *float64    `json:"AccMinZ"`
	GyrMeanX      *float64    `json:"GyrMeanX"`
	GyrMeanY      *float64    `json:"GyrMeanY"`
	GyrMeanZ      *float64    `json:"GyrMeanZ"`
	GyrVarianceX  *float64    `json:"GyrVarX"`
	GyrVarianceY  *float64    `json:"GyrVarY"`
	GyrVarianceZ  *float64    `json:"GyrVarZ"`
	GyrMedianX    *float64    `json:"GyrMedX"`
	GyrMedianY    *float64    `json:"GyrMedY"`
	GyrMedianZ    *float64    `json:"GyrMedZ"`
	GyrStdDevX    *float64    `json:"GyrStdX"`
	GyrStdDevY    *float64    `json:"GyrStdY"`
	GyrStdDevZ    *float64    `json:"GyrStdZ"`
	GyrSkewX      *float64    `json:"GyrSkewX"`
	GyrSkewY      *float64    `json:"GyrSkewY"`
	GyrSkewZ      *float64    `json:"GyrSkewZ"`
	GyrMaxX       *float64    `json:"GyrMaxX"`
	GyrMaxY       *float64    `json:"GyrMaxY"`
	GyrMaxZ       *float64    `json:"GyrMaxZ"`
	GyrMinX       *float64    `json:"GyrMinX"`
	GyrMinY       *float64    `json:"GyrMinY"`
	GyrMinZ       *float64    `json:"GyrMinZ"`
	RollMean      *float64    `json:"RollMean"`
	RollVariance  *float64    `json:"RollVar"`
	RollMedian    *float64    `json:"RollMed"`
	RollStdDev    *float64    `json:"RollStd"`
	RollSkew      *float64    `json:"RollSkew"`
	RollMax       *float64    `json:"RollMax"`
	RollMin       *float64    `json:"RollMin"`
	PitchMean     *float64    `json:"PitchMean"`
	PitchVariance *float64    `json:"PitchVar"`
	PitchMedian   *float64    `json:"PitchMed"`
	PitchStdDev   *float64    `json:"PitchStd"`
	PitchSkew     *float64    `json:"PitchSkew"`
	PitchMax      *float64    `json:"PitchMax"`
	PitchMin      *float64    `json:"PitchMin"`
	GestureID     *int        `json:"GestureID,omitempty"`
	DataPoints    []DataPoint `json:"DataPoints"`
}

// ValidateFeatureDataNotNil checks if any field in FeatureData is nil
func ValidateFeatureDataNotNil(featureData FeatureData) error {
	fields := map[string]*float64{
		"AccMeanX":      featureData.AccMeanX,
		"AccMeanY":      featureData.AccMeanY,
		"AccMeanZ":      featureData.AccMeanZ,
		"AccVarianceX":  featureData.AccVarianceX,
		"AccVarianceY":  featureData.AccVarianceY,
		"AccVarianceZ":  featureData.AccVarianceZ,
		"AccMedianX":    featureData.AccMedianX,
		"AccMedianY":    featureData.AccMedianY,
		"AccMedianZ":    featureData.AccMedianZ,
		"AccStdDevX":    featureData.AccStdDevX,
		"AccStdDevY":    featureData.AccStdDevY,
		"AccStdDevZ":    featureData.AccStdDevZ,
		"AccSkewX":      featureData.AccSkewX,
		"AccSkewY":      featureData.AccSkewY,
		"AccSkewZ":      featureData.AccSkewZ,
		"AccMaxX":       featureData.AccMaxX,
		"AccMaxY":       featureData.AccMaxY,
		"AccMaxZ":       featureData.AccMaxZ,
		"AccMinX":       featureData.AccMinX,
		"AccMinY":       featureData.AccMinY,
		"AccMinZ":       featureData.AccMinZ,
		"GyrMeanX":      featureData.GyrMeanX,
		"GyrMeanY":      featureData.GyrMeanY,
		"GyrMeanZ":      featureData.GyrMeanZ,
		"GyrVarianceX":  featureData.GyrVarianceX,
		"GyrVarianceY":  featureData.GyrVarianceY,
		"GyrVarianceZ":  featureData.GyrVarianceZ,
		"GyrMedianX":    featureData.GyrMedianX,
		"GyrMedianY":    featureData.GyrMedianY,
		"GyrMedianZ":    featureData.GyrMedianZ,
		"GyrStdDevX":    featureData.GyrStdDevX,
		"GyrStdDevY":    featureData.GyrStdDevY,
		"GyrStdDevZ":    featureData.GyrStdDevZ,
		"GyrSkewX":      featureData.GyrSkewX,
		"GyrSkewY":      featureData.GyrSkewY,
		"GyrSkewZ":      featureData.GyrSkewZ,
		"GyrMaxX":       featureData.GyrMaxX,
		"GyrMaxY":       featureData.GyrMaxY,
		"GyrMaxZ":       featureData.GyrMaxZ,
		"GyrMinX":       featureData.GyrMinX,
		"GyrMinY":       featureData.GyrMinY,
		"GyrMinZ":       featureData.GyrMinZ,
		"RollMean":      featureData.RollMean,
		"RollVariance":  featureData.RollVariance,
		"RollMedian":    featureData.RollMedian,
		"RollStdDev":    featureData.RollStdDev,
		"RollSkew":      featureData.RollSkew,
		"RollMax":       featureData.RollMax,
		"RollMin":       featureData.RollMin,
		"PitchMean":     featureData.PitchMean,
		"PitchVariance": featureData.PitchVariance,
		"PitchMedian":   featureData.PitchMedian,
		"PitchStdDev":   featureData.PitchStdDev,
		"PitchSkew":     featureData.PitchSkew,
		"PitchMax":      featureData.PitchMax,
		"PitchMin":      featureData.PitchMin,
	}

	for fieldName, fieldValue := range fields {
		if fieldValue == nil {
			return errors.New("field " + fieldName + " cannot be nil")
		}
	}

	return nil
}

// CreateFeatureTable creates the feature_data table if it does not exist
func CreateFeatureTable(tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS feature_data (
		id SERIAL PRIMARY KEY,
		acc_mean_x FLOAT NOT NULL,
		acc_mean_y FLOAT NOT NULL,
		acc_mean_z FLOAT NOT NULL,
		acc_variance_x FLOAT NOT NULL,
		acc_variance_y FLOAT NOT NULL,
		acc_variance_z FLOAT NOT NULL,
		acc_median_x FLOAT NOT NULL,
		acc_median_y FLOAT NOT NULL,
		acc_median_z FLOAT NOT NULL,
		acc_std_dev_x FLOAT NOT NULL,
		acc_std_dev_y FLOAT NOT NULL,
		acc_std_dev_z FLOAT NOT NULL,
		acc_skew_x FLOAT NOT NULL,
		acc_skew_y FLOAT NOT NULL,
		acc_skew_z FLOAT NOT NULL,
		acc_max_x FLOAT NOT NULL,
		acc_max_y FLOAT NOT NULL,
		acc_max_z FLOAT NOT NULL,
		acc_min_x FLOAT NOT NULL,
		acc_min_y FLOAT NOT NULL,
		acc_min_z FLOAT NOT NULL,
		gyr_mean_x FLOAT NOT NULL,
		gyr_mean_y FLOAT NOT NULL,
		gyr_mean_z FLOAT NOT NULL,
		gyr_variance_x FLOAT NOT NULL,
		gyr_variance_y FLOAT NOT NULL,
		gyr_variance_z FLOAT NOT NULL,
		gyr_median_x FLOAT NOT NULL,
		gyr_median_y FLOAT NOT NULL,
		gyr_median_z FLOAT NOT NULL,
		gyr_std_dev_x FLOAT NOT NULL,
		gyr_std_dev_y FLOAT NOT NULL,
		gyr_std_dev_z FLOAT NOT NULL,
		gyr_skew_x FLOAT NOT NULL,
		gyr_skew_y FLOAT NOT NULL,
		gyr_skew_z FLOAT NOT NULL,
		gyr_max_x FLOAT NOT NULL,
		gyr_max_y FLOAT NOT NULL,
		gyr_max_z FLOAT NOT NULL,
		gyr_min_x FLOAT NOT NULL,
		gyr_min_y FLOAT NOT NULL,
		gyr_min_z FLOAT NOT NULL,
		roll_mean FLOAT NOT NULL,
		roll_variance FLOAT NOT NULL,
		roll_median FLOAT NOT NULL,
		roll_std_dev FLOAT NOT NULL,
		roll_skew FLOAT NOT NULL,
		roll_max FLOAT NOT NULL,
		roll_min FLOAT NOT NULL,
		pitch_mean FLOAT NOT NULL,
		pitch_variance FLOAT NOT NULL,
		pitch_median FLOAT NOT NULL,
		pitch_std_dev FLOAT NOT NULL,
		pitch_skew FLOAT NOT NULL,
		pitch_max FLOAT NOT NULL,
		pitch_min FLOAT NOT NULL,
		gestureid INT REFERENCES gestures(id) ON DELETE CASCADE
	)`
	_, err := tx.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("Error creating feature_data table: %v\n", err)
	}
	log.Printf("Created Feature table successfully")
	return nil
}

func InsertFeatureData(tx pgx.Tx, featureData *FeatureData, gestureID int) error {
	var data = *featureData
	query := `INSERT INTO feature_data (
		acc_mean_x, acc_mean_y, acc_mean_z, 
		acc_variance_x, acc_variance_y, acc_variance_z,
		acc_median_x, acc_median_y, acc_median_z, 
		acc_std_dev_x, acc_std_dev_y, acc_std_dev_z,
		acc_skew_x, acc_skew_y, acc_skew_z, 
		acc_max_x, acc_max_y, acc_max_z, 
		acc_min_x, acc_min_y, acc_min_z,
		gyr_mean_x, gyr_mean_y, gyr_mean_z, 
		gyr_variance_x, gyr_variance_y, gyr_variance_z,
		gyr_median_x, gyr_median_y, gyr_median_z, 
		gyr_std_dev_x, gyr_std_dev_y, gyr_std_dev_z,
		gyr_skew_x, gyr_skew_y, gyr_skew_z, 
		gyr_max_x, gyr_max_y, gyr_max_z, 
		gyr_min_x, gyr_min_y, gyr_min_z,
		roll_mean, roll_variance, roll_median, roll_std_dev, roll_skew, roll_max, roll_min,
		pitch_mean, pitch_variance, pitch_median, pitch_std_dev, pitch_skew, pitch_max, pitch_min,
		gestureid
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21,
		$22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42,
		$43, $44, $45, $46, $47, $48, $49,
		$50, $51, $52, $53, $54, $55, $56,
		$57
	) RETURNING id`

	err := tx.QueryRow(context.Background(), query,
		data.AccMeanX, data.AccMeanY, data.AccMeanZ, data.AccVarianceX, data.AccVarianceY, data.AccVarianceZ,
		data.AccMedianX, data.AccMedianY, data.AccMedianZ, data.AccStdDevX, data.AccStdDevY, data.AccStdDevZ,
		data.AccSkewX, data.AccSkewY, data.AccSkewZ, data.AccMaxX, data.AccMaxY, data.AccMaxZ, data.AccMinX, data.AccMinY, data.AccMinZ,
		data.GyrMeanX, data.GyrMeanY, data.GyrMeanZ, data.GyrVarianceX, data.GyrVarianceY, data.GyrVarianceZ,
		data.GyrMedianX, data.GyrMedianY, data.GyrMedianZ, data.GyrStdDevX, data.GyrStdDevY, data.GyrStdDevZ,
		data.GyrSkewX, data.GyrSkewY, data.GyrSkewZ, data.GyrMaxX, data.GyrMaxY, data.GyrMaxZ, data.GyrMinX, data.GyrMinY, data.GyrMinZ,
		data.RollMean, data.RollVariance, data.RollMedian, data.RollStdDev, data.RollSkew, data.RollMax, data.RollMin,
		data.PitchMean, data.PitchVariance, data.PitchMedian, data.PitchStdDev, data.PitchSkew, data.PitchMax, data.PitchMin,
		gestureID,
	).Scan(&featureData.ID)

	if err != nil {
		return fmt.Errorf("Error storing feature data: %v\n", err)
	}
	return nil
}

// GetFeatureDataByID retrieves feature data by ID
func GetFeatureDataByID(id *int) (*FeatureData, error) {
	if id == nil {
		return nil, fmt.Errorf("ID is nil")
	}
	conn := GetConn()
	if conn == nil {
		return nil, fmt.Errorf("Database connection is nil")
	}

	var data FeatureData
	query := `SELECT 
		id, acc_mean_x, acc_mean_y, acc_mean_z, 
		acc_variance_x, acc_variance_y, acc_variance_z,
		acc_median_x, acc_median_y, acc_median_z, 
		acc_std_dev_x, acc_std_dev_y, acc_std_dev_z,
		acc_skew_x, acc_skew_y, acc_skew_z, 
		acc_max_x, acc_max_y, acc_max_z, 
		acc_min_x, acc_min_y, acc_min_z,
		gyr_mean_x, gyr_mean_y, gyr_mean_z, 
		gyr_variance_x, gyr_variance_y, gyr_variance_z,
		gyr_median_x, gyr_median_y, gyr_median_z, 
		gyr_std_dev_x, gyr_std_dev_y, gyr_std_dev_z,
		gyr_skew_x, gyr_skew_y, gyr_skew_z, 
		gyr_max_x, gyr_max_y, gyr_max_z, 
		gyr_min_x, gyr_min_y, gyr_min_z,
		roll_mean, roll_variance, roll_median, roll_std_dev, roll_skew, roll_max, roll_min,
		pitch_mean, pitch_variance, pitch_median, pitch_std_dev, pitch_skew, pitch_max, pitch_min,
		gestureid
	FROM feature_data WHERE id = $1`

	err := conn.QueryRow(context.Background(), query, id).Scan(
		&data.ID, &data.AccMeanX, &data.AccMeanY, &data.AccMeanZ, &data.AccVarianceX, &data.AccVarianceY, &data.AccVarianceZ,
		&data.AccMedianX, &data.AccMedianY, &data.AccMedianZ, &data.AccStdDevX, &data.AccStdDevY, &data.AccStdDevZ,
		&data.AccSkewX, &data.AccSkewY, &data.AccSkewZ, &data.AccMaxX, &data.AccMaxY, &data.AccMaxZ, &data.AccMinX, &data.AccMinY, &data.AccMinZ,
		&data.GyrMeanX, &data.GyrMeanY, &data.GyrMeanZ, &data.GyrVarianceX, &data.GyrVarianceY, &data.GyrVarianceZ,
		&data.GyrMedianX, &data.GyrMedianY, &data.GyrMedianZ, &data.GyrStdDevX, &data.GyrStdDevY, &data.GyrStdDevZ,
		&data.GyrSkewX, &data.GyrSkewY, &data.GyrSkewZ, &data.GyrMaxX, &data.GyrMaxY, &data.GyrMaxZ, &data.GyrMinX, &data.GyrMinY, &data.GyrMinZ,
		&data.RollMean, &data.RollVariance, &data.RollMedian, &data.RollStdDev, &data.RollSkew, &data.RollMax, &data.RollMin,
		&data.PitchMean, &data.PitchVariance, &data.PitchMedian, &data.PitchStdDev, &data.PitchSkew, &data.PitchMax, &data.PitchMin,
		&data.GestureID,
	)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving feature data: %v", err)
	}
	return &data, nil
}

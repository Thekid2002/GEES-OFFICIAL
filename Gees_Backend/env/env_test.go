package env

import (
	"Gees_Backend/models"
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	envString := "test.env"
	LoadEnv(&envString)
	err := models.Connect()
	if err != nil {
		panic(err.Error())
	}
	code := m.Run()

	tx, err := models.StartTransaction()
	if err != nil {
		panic(err.Error())
	}

	_, err = tx.Exec(context.Background(), `DROP TABLE IF EXISTS gesture_action_mapping CASCADE; DROP TABLE IF EXISTS feature_data CASCADE; DROP TABLE IF EXISTS gestures CASCADE; DROP TABLE IF EXISTS data_points CASCADE; DROP TABLE IF EXISTS actions CASCADE`)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err != nil {
			err = models.RollbackTransaction(tx)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	err = models.CommitTransaction(tx)
	if err != nil {
		panic(err.Error())
	}
	os.Exit(code)
}

package env

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv(nil)

	if os.Getenv("DATABASE_URL") == "" {
		t.Error("DATABASE_URL is not set")
	}

	if os.Getenv("PYTHON_URL") == "" {
		t.Error("PYTHON_URL is not set")
	}

	if os.Getenv("ENV") == "" {
		t.Error("ENV is not set")
	}
}

package env

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

func getBaseDir() string {
	_, b, _, _ := runtime.Caller(0) // Get the current file path
	return filepath.Dir(b)          // Return the directory of the current file
}

func LoadEnv(forcedEnv *string) {
	var envFile string

	if forcedEnv != nil {
		envFile = *forcedEnv
	}

	if forcedEnv == nil {
		env := flag.String("env", "development", "Environment to load (development, docker)")
		flag.Parse()

		switch *env {
		case "docker":
			envFile = "docker.env"
		default:
			envFile = "development.env"
		}
	}

	baseDir := getBaseDir()
	envFile = filepath.Join(baseDir, envFile) // Use baseDir to locate the env folder
	absPath, err := filepath.Abs(envFile)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	fmt.Printf("Loading environment variables from .env file %s\n", absPath)
	err = godotenv.Load(absPath)
	if err != nil {
		log.Printf("Warning: Could not load %s file: %v", absPath, err)
	}
}

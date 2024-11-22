package initializers

import (
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

var PORT string

func LoadEnvVariables() {

	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
	}

	// Construct the absolute path to the .env file
	envPath := filepath.Join(workingDir, ".env")
	// Load the .env file
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file at path:", envPath, "Error:", err)
	}
}

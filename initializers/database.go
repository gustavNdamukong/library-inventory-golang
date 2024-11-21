package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//var DBConnection *gorm.DB

func DBConnection() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_CREDENTIALS")
	fmt.Println("Using DSN:", dsn) // Add this line
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		log.Println("Connected to database")
	}

	log.Fatal("STOPPING HERE")
	return DB
}

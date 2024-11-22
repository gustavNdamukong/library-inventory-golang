package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_CREDENTIALS")
	log.Println("GOT DSN: ", dsn)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		log.Println("Connected to database")
	}

	return DB
}

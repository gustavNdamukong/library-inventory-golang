package main

import (
	"github.com/gustavNdamukong/library-inventory-golang/initializers"
	"github.com/gustavNdamukong/library-inventory-golang/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnvVariables()
	DB = initializers.ConnectToDB()
}

func main() {
	DB.AutoMigrate(&models.Books{})
}

/*
	NOTES:
	To run migrations in Go with GORM, follow these steps:
	-create a model package on your root in a 'models' dir
	-create models inside this dir eg a book model file will be named 'booksModel.go'.
		Here are its contents:

			package models

			import "gorm.io/gorm"

			type Books struct {
				gorm.Model
				ID       string
				Title    string
				Author   string
				Quantity int
				Stock    int
			}
	-Then create a migration package on ur root in a 'migrations' dir
	-Create 'migrate.go' file in there, and place this code in it:

			package main

			import (
				"github.com/gustavNdamukong/library-inventory-golang/initializers"
				"github.com/gustavNdamukong/library-inventory-golang/models"
				"gorm.io/gorm"
			)

			var DB *gorm.DB

			func init() {
				initializers.LoadEnvVariables()
				DB = initializers.ConnectToDB()
			}

			func main() {
				DB.AutoMigrate(&models.Books{})
			}

	-That's it. To run your migration, juast run this command:

			go run migrations/migrate.go

	-The table books will be created as defined in the Books struct in booksModel.go
		and any data in it will be inserted as well.
*/

package main

import (
	"fmt"
	"os"

	"github.com/gustavNdamukong/library-inventory-golang/controllers"
	"github.com/gustavNdamukong/library-inventory-golang/initializers"

	"github.com/gin-gonic/gin"
)

var PORT string

func init() {
	initializers.LoadEnvVariables()

	PORT = os.Getenv("PORT")
}

func main() {

	//create a router for your app
	router := gin.Default()
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.BookById)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/checkout", controllers.CheckoutBook)
	router.PATCH("/books/return", controllers.ReturnBook)

	//run your new Gin web server on any port of your choice
	router.Run(fmt.Sprintf("localhost:%s", PORT))
	//the end point of the API is 'http://localhost:8080/books'
	//to run your app & expose the server, you should do 'go run main.go' from within your app root.
	//you would get an output like: "Listening and serving HTTP on localhost:8080"
}

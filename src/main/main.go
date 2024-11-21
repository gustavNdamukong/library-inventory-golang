package main

import (
	"fmt"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

const port = 8080

type book struct {
	//the first letters of these keys MUST be uppercase. This makes the fields publicly accessible
	// fields by other modules outside of this (main) module we are in.
	//Cast it (so to speak) to a json equalvalent as we'll need the json version in the API calls
	//Stock is the record of the library's original inventory for each book

	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Stock    int    `json:"stock"`
}

// This book data would normally come from a DB. We will use a variable to store them in memory.
// Notice that trhis books array gets its data from a serialised version of the book struct above.
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2, Stock: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5, Stock: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6, Stock: 6},
}

func init() {
	/*err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}*/
	initializers.LoadEnvVariables()
}

// a gin context contains all the info about a request. It allows u return a response
func getBooks(c *gin.Context) {
	//specify that you wish to get back a nicely formatted json obj (with proper indentation)
	//we send back a status of 200 (http.StatusOK) & the books data (wh will be auto serialised 4 us into json)
	//We could have returned files, HTML pages etc
	c.JSON(200, gin.H{
		"message": "Yeeeeaaahhh - books right?",
	})
	///// c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	//check if we have an 'id' as a query string parameter & return an error if not
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	//Check if this library keeps that book & return an error if not
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	//if this library keeps that book, check if it's currently available (quantity is more than 0)
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	//if available, reduce the current quantity by one to check out one
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// Allow people to return checked-out books
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	//Check if this library keeps that book & return an error if not
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	//if this library keeps that book, increment its quantity by one to reflect that the borrowed book was returned
	//until all its original inventary stock has been returned
	if book.Quantity < book.Stock {
		book.Quantity += 1
		c.IndentedJSON(http.StatusOK, book)
		return
	}

	//let them know that this no further returns are expected for this book.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sorry! No further returns are expected for this book."})
}

// allow for the creation of a book resource
func createBook(c *gin.Context) {
	var newBook book

	//bind the data sent in this (current) request to a variable (in this case newbook)
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	//send back a response with status code 201 (http.StatusCreated)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {

	//create a router for your app
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)

	//run your new Gin web server on any port of your choice
	router.Run(fmt.Sprintf("localhost:%d", port))
	//After this line, we've just implemented an API (albeit it has only one func for now-getBookd())
	//the end point of the API is 'http://localhost:8080/books'
	//to run your app & expose the server, you should do 'go run main.go' from within your src/main dir.
	//you would get an output like: "Listening and serving HTTP on localhost:8080"
}

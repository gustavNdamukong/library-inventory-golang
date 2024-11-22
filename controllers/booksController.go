package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavNdamukong/library-inventory-golang/initializers"
	"github.com/gustavNdamukong/library-inventory-golang/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnvVariables()
	DB = initializers.ConnectToDB()
}

// allow for the creation of a book resource
func CreateBook(c *gin.Context) {
	/* NOTES: Gorm models. gorm.Model comes with three fields pre-defined:
	-CreatedAt
	-UpdatedAt
	-DeletedAt

	So you do not have to have these fields on your model. The only thing yo MUST do is
	make sure your model table has the three fields 'created_at', 'updated_at', & 'deleted_at"
	with timestamps of 'timestamptz'.
	-Name your model files with a prefix of the table name in lowercase, followed by 'Model'.
	 For example, we name the model of the books DB table 'booksModel'.
	-In the model file, which is oftern in a models package in 'models/', we declare the model
	 as a struct which matches the DB table name beginning in uppercase. For example, the model
	of the books table in 'models/booksModel.go' has the following code:

		type Books struct {
			gorm.Model
			ID       int    `json:"id"`
			Title    string `json:"title"`
			Author   string `json:"author"`
			Quantity int    `json:"quantity"`
			Stock    int    `json:"stock"`
		}

	*/
	var newBook *models.Books

	// get the data sent in this (current) request, and bind it to a variable (newbook)
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// create a book record
	book := models.Books{
		Title:    newBook.Title,
		Author:   newBook.Author,
		Quantity: newBook.Quantity,
		Stock:    newBook.Stock,
	}

	result := DB.Create(&book)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//send back a response with status code 201 (http.StatusCreated) & the newly created book data
	c.IndentedJSON(http.StatusCreated, newBook)
}

// a gin context contains all the info about a request. It allows u return a response
func GetBooks(c *gin.Context) {
	var booksModel []models.Books
	books := DB.Find(&booksModel)

	if books.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": books.Error.Error(),
		})
		return
	}

	//specify that you wish to get back a nicely formatted json obj (with proper indentation)
	//we send back a status of 200 (http.StatusOK) & the books data (wh will be auto serialised 4 us into json)
	//We could have returned files, HTML pages etc

	/*c.JSON(200, gin.H{
		"message": books,
	})*/

	c.IndentedJSON(http.StatusOK, booksModel)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func GetBookById(id string) (*models.Books, error) {
	// similar to GetPosts, only here, bookModel is not an array
	var bookModel models.Books
	result := DB.First(&bookModel, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &bookModel, nil
}

func CheckoutBook(c *gin.Context) {
	//check if we have an 'id' as a query string parameter & return an error if not
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	//Check if this library keeps that book & return an error if not
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	//if this library keeps that book, check if it's currently available (quantity is more than 0)
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	// update the record here
	//if available, reduce the current quantity by one to check out one
	/*
		NOTES: How to do an update in GORM.
			db.First(&user)

			user.Name = "jinzhu 2"
			user.Age = 100
			db.Save(&user)

			So
				db.Save(&User{Name: "jinzhu", Age: 100})

			will do an INSERT
			While
				db.Save(&User{ID: 1, Name: "jinzhu", Age: 100})
			will do an UPDATE
	*/
	book.Quantity -= 1
	DB.Save(&book)

	c.IndentedJSON(http.StatusOK, book)
}

// Allow people to return checked-out books
func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	//Check if this library keeps that book & return an error if not
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	//if this library keeps that book, increment its quantity by one to reflect that the borrowed book was returned
	//until all its original inventary stock has been returned
	if book.Quantity < book.Stock {
		book.Quantity += 1
		// update the record here
		DB.Save(&book)
		c.IndentedJSON(http.StatusOK, book)
		return
	}

	//let them know that no further returns are expected for this book.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sorry! No further returns are expected for this book."})
}

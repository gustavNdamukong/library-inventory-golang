package controllers

import "github.com/gin-gonic/gin"

// a gin context contains all the info about a request. It allows u return a response
func GetBooks(c *gin.Context) {
	//specify that you wish to get back a nicely formatted json obj (with proper indentation)
	//we send back a status of 200 (http.StatusOK) & the books data (wh will be auto serialised 4 us into json)
	//We could have returned files, HTML pages etc
	c.JSON(200, gin.H{
		"message": "Yeeeeaaahhh - books right? we are in the controller",
	})
	///// c.IndentedJSON(http.StatusOK, books)
}

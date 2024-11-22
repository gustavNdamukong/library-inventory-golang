# A library book inventrory management API built in Golang and Gin

## Meant to demonstrate the use of the Gin framework with DB access using the GIN ORM GORM

* Features
    * View all books
    * Fetch a book by its ID
    * Create a book
    * Checkout a book
    * Return a book
    * Good error handling for cases where a book is not found, or a book is no 
        longer available (all checked out), or when no further returns are needed,
        because, all checked out books on a particular ID have been returned. 

## Demonstrates the power of GORM
* How to work with models
    * How to use controllers as request handlers
    * Handling and responding to requests based on various HTTP headers
    * Manage environmental variables using the godotenv package
* DB operations in GORM
    * PostgreSQL database connection
    * Learn how to create a new record (INSERT)
    * Learn how to fetch all records
    * Learn how to fetch a singe record by its ID
    * Learn how to update a record by its ID

# Info
To get your started, there is an SQL script in the 'books.sql' file which you 
can use to create the books table, and insert some test data in it. 
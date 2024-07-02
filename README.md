# A library book inventrory management API build in Golang and Gin

## Meant to demonstrate the use of the Gin frameword

* Features
    * View all books
    * Fetch a book by its ID
    * Create a book
    * Checkout a book
    * Return a book
    * Good error handling for cases where a book is not found, or a book is no 
        longer available (all checked out), and when no further returns are needed,
        as in, when all checked out books on a particular ID have all been returned. 

For now, this app does not pull from any database. Its book records are stored 
in memory (in a variable), so its sd not persiste across server restarts.
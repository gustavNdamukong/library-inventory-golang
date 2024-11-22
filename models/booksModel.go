package models

import (
	"gorm.io/gorm"
)

// the first letters of these keys MUST be uppercase. This makes the fields publicly accessible
// fields by other modules outside of this (main) module we are in.
// Cast it (so to speak) to a json equalvalent as we'll need the json version in the API calls
// Stock is the record of the library's original inventory for each book
type Books struct {
	gorm.Model
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Stock    int    `json:"stock"`
}

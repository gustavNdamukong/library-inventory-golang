package models

import "gorm.io/gorm"

type Bindas struct {
	gorm.Model
	ID       string
	Title    string
	Author   string
	Quantity int
	Stock    int
}

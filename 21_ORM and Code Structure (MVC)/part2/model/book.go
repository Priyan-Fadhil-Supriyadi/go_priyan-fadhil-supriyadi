package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Author string `json:"author" form:"author"`
	Title  string `json:"title" form:"title"`
	Price  int    `json:"price" form:"price"`
}

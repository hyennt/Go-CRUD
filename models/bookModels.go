package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string `json:"title"`
	AuthorId uint   `json:"author_id"`
}

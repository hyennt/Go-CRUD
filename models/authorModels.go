package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `json:"books" gorm:"foreignKey:AuthorId"`
}

type Book struct {
	gorm.Model
	Title    string `json:"title"`
	AuthorId uint   `json:"author_id"`
}

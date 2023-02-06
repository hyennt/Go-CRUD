package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string
	Body     string
	AuthorId string
	Books    []Book `json:"books" gorm:"foreignKey:PostId"`
}

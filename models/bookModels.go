package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title      string `json:"title"`
	AuthorId   uint   `json:"author_id"`
	CategoryId uint   `json:"category_id"`
	PostId     uint   `json:"post_id"`
	//Category   Category `json:"category" gorm:"foreignKey:CategoryId"`
}

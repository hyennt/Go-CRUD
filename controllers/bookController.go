package controllers

import (
	"go-crud/initialize"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func BookGetting(c *gin.Context) {
	var books []models.Book
	initialize.DB.Find(&books)
	c.IndentedJSON(200, gin.H{
		"books": books,
	})
}

func BookCreate(c *gin.Context) {
	var body struct {
		Title      string `json:"title"`
		AuthorId   uint   `json:"author_id"`
		CategoryId uint   `json:"category_id"`
		PostId     uint   `json:"post_id"`
	}
	c.BindJSON(&body)
	books := models.Book{Title: body.Title, AuthorId: body.AuthorId, CategoryId: body.CategoryId, PostId: body.PostId}
	result := initialize.DB.Create(&books)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.IndentedJSON(200, gin.H{
		"book":    books,
		"message": "Book created successfully",
	})
}

func BookShowByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initialize.DB.First(&book, id)
	c.IndentedJSON(200, gin.H{
		"data": []gin.H{
			{
				"book": book,
			},
		},
		"links": []gin.H{
			{
				"method":      "GET",
				"author_path": "/api/author/author_id",
				"author_id":   book.AuthorId,
			},
		},
		"links_category": []gin.H{
			{
				"method":        "GET",
				"category_path": "/api/category/category_id",
				"category_id":   book.CategoryId,
			},
		},
	})
}

func BookDelete(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initialize.DB.First(&book, id)
	initialize.DB.Delete(&book)
	c.IndentedJSON(200, gin.H{
		"message": "Book deleted successfully",
	})
}

func BookUpdate(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initialize.DB.First(&book, id)
	var body struct {
		Title      string          `json:"title"`
		AuthorId   uint            `json:"author_id"`
		CategoryId uint            `json:"category_id"`
		Category   models.Category `json:"category"`
	}
	c.BindJSON(&body)
	initialize.DB.Model(&book).Updates(models.Book{Title: body.Title, AuthorId: body.AuthorId})
	c.IndentedJSON(200, gin.H{
		"message":     "Book updated successfully",
		"bookUpdated": book,
	})
}

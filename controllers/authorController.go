package controllers

import (
	"go-crud/initialize"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func AuthorGetting(c *gin.Context) {
	var authors []models.Author
	initialize.DB.Find(&authors)
	c.IndentedJSON(200, gin.H{
		"authors": authors,
	})
}

func BookGetting(c *gin.Context) {
	var books []models.Book
	initialize.DB.Find(&books)
	c.IndentedJSON(200, gin.H{
		"books": books,
	})
}

func BookCreate(c *gin.Context) {
	var body struct {
		Title    string `json:"title"`
		AuthorId uint   `json:"author_id"`
	}
	c.BindJSON(&body)
	books := models.Book{Title: body.Title, AuthorId: body.AuthorId}
	result := initialize.DB.Create(&books)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.IndentedJSON(200, gin.H{
		"message": "Book created successfully",
	})
}

func AuthorCreate(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Books []struct {
			Title string `json:"title"`
		} `json:"books"`
	}
	c.BindJSON(&body)
	author := models.Author{Name: body.Name, Books: []models.Book{}}
	result := initialize.DB.Create(&author)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.IndentedJSON(200, gin.H{
		"message": "Author created successfully",
		"hateoas": []gin.H{
			{"method": "GET", "href": "/api/author"},
		},
	})
}

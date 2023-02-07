package controllers

import (
	"go-crud/initialize"
	"go-crud/models"
	"strconv"

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
	//authorId := c.Param("authorId")
	//categoryId := c.Param("categoryId")
	var book models.Book
	var author models.Author
	var category models.Category
	initialize.DB.First(&book, id)
	initialize.DB.First(&author, book.AuthorId)
	author_ID := strconv.Itoa(int(book.AuthorId))
	initialize.DB.First(&category, book.CategoryId)
	category_ID := strconv.Itoa(int(book.CategoryId))
	c.IndentedJSON(200, gin.H{
		"data": []gin.H{
			{
				"book": book,
			},
		},

		"links": []gin.H{
			{
				"self":   "http://localhost:3000/api/book/" + id,
				"method": "GET",
			},
		},

		"attributed_by": []gin.H{
			{
				"author_info":      author,
				"method":           "GET",
				"author_self_link": "http://localhost:3000/api/author/" + author_ID,
			},
		},

		"Related_to": []gin.H{
			{
				"category_info":      category,
				"method":             "GET",
				"category_self_link": "http://localhost:3000/api/category/" + category_ID,
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

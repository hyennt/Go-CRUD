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

func AuthorShowByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	initialize.DB.First(&author, id)

	c.IndentedJSON(200, gin.H{
		"author": author,
	})
}

func AuthorUpdate(c *gin.Context) {
	id := c.Param("id")

	var author models.Author

	var body struct {
		Name  string `json:"name"`
		Books []struct {
			Title string `json:"title"`
		} `json:"books"`
	}
	c.BindJSON(&body)

	//find post with id
	initialize.DB.First(&author, id)

	//update
	initialize.DB.Model(&author).Updates(models.Author{Name: body.Name, Books: []models.Book{}})
	c.IndentedJSON(200, gin.H{
		"author":  author,
		"message": "Author updated successfully",
	})

}

func AuthorDelete(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	initialize.DB.First(&author, id)
	initialize.DB.Delete(&author)
	c.IndentedJSON(200, gin.H{
		"message": "Author deleted successfully",
	})
}

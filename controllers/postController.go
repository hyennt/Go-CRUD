package controllers

import (
	"go-crud/initialize"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Find post with id
	// id := c.Param("id")
	// var post2 models.Post
	// initialize.DB.First(&post2, id)

	// Get data off body req
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	c.BindJSON(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initialize.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.IndentedJSON(200, gin.H{
		"message": "Post created successfully",
		"hateoas": []gin.H{
			{"method": "GET", "href": "/getPost"},
			{"method": "GET", "href": "/getPost/:id"},
			{"method": "PUT", "href": "/update/:id"},
			{"method": "DELETE", "href": "/delete/:id"},
		},
	})
}

func PostGetting(c *gin.Context) {
	var posts []models.Post
	initialize.DB.Find(&posts)
	c.IndentedJSON(200, gin.H{
		"posts": posts,
	})
}

func PostShowByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	// find post with id
	initialize.DB.First(&post, id)

	c.IndentedJSON(200, gin.H{
		"post": post,
		"hateoas": []gin.H{
			{"method": "GET", "href": "/getPost"},
			{"method": "PUT", "href": "/update/:id"},
			{"method": "DELETE", "href": "/delete/:id"},
		},
		"base": "http://localhost:3000/getPost/3",
	})
	//fmt.Println("PostShowByID", id)
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	c.BindJSON(&body)

	//find post with id
	initialize.DB.First(&post, id)

	//update post
	initialize.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.IndentedJSON(200, gin.H{
		"post": post,
		"hateoas": []gin.H{
			{"rel": "posts", "href": "/getPost"},
			{"method": "GET", "href": "/getPost/:id"},
			{"method": "DELETE", "href": "/delete/:id"},
		},
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	//find post with id
	initialize.DB.First(&post, id)

	//delete post
	initialize.DB.Delete(&post, id)

	c.IndentedJSON(200, gin.H{
		"message": "Post deleted successfully",
		"hateoas": []gin.H{
			{"rel": "posts", "href": "/getPost"},
			{"method": "GET", "href": "/getPost/:id"},
			{"method": "PUT", "href": "/update/:id"},
		},
	})
}

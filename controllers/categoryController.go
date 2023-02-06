package controllers

import (
	"go-crud/initialize"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func CategoryGetting(c *gin.Context) {
	var categories []models.Category
	initialize.DB.Find(&categories)
	c.IndentedJSON(200, gin.H{
		"categories": categories,
	})
}

func CategoryCreate(c *gin.Context) {
	var body struct {
		Type string `json:"type"`
	}

	c.BindJSON(&body)

	categories := models.Category{Type: body.Type}
	result := initialize.DB.Create(&categories)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.IndentedJSON(200, gin.H{
		"category": categories,
		"message":  "Category created successfully",
	})
}

func CategoryShowByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	initialize.DB.First(&category, id)
	c.IndentedJSON(200, gin.H{
		"category": category,
	})
}

func CategoryDelete(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	initialize.DB.First(&category, id)
	initialize.DB.Delete(&category)
	c.IndentedJSON(200, gin.H{
		"message": "Category deleted successfully",
	})
}

func CategoryUpdate(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	initialize.DB.First(&category, id)
	var body struct {
		Type string `json:"type"`
	}
	c.BindJSON(&body)
	category.Type = body.Type
	initialize.DB.Save(&category)
	c.IndentedJSON(200, gin.H{
		"message": "Category updated successfully",
	})
}

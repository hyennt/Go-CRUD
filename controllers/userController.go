package controllers

import (
	"go-crud/initialize"
	"go-crud/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Failed to ready body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	// Hashing the password with the default cost of 10
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	users := models.User{Email: body.Email, Password: string(hash)}
	result := initialize.DB.Create(&users)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"message": "User created successfully",
	})

}

func UserDelete(c *gin.Context) {
	//id := c.Param("id")
	var user models.User
	// DELETE ALL

	initialize.DB.Delete(&user)
	//initialize.DB.Delete(&user)
	c.IndentedJSON(200, gin.H{
		"message": "User deleted successfully",
	})
}

func Login(c *gin.Context) {
	// Get the email and password off the request body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Failed to ready body",
		})
		return
	}
	// Look up requested user in the database

	var user models.User
	initialize.DB.Where(&user, "email = ?", body.Email).Find(&user)

	// Compare the password with the hashed password in the database

	// Generate a JWT token and send it back to the client
}

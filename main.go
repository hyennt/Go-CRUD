package main

import (
	"go-crud/controllers"
	"go-crud/initialize"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnvVar()
	initialize.ConnectDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/getPost", controllers.PostGetting)
	r.GET("/getPost/:id", controllers.PostShowByID)
	r.PUT("/update/:id", controllers.PostUpdate)
	r.DELETE("/delete/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}

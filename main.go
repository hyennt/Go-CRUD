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
	// r.POST("/posts", controllers.PostCreate)
	// r.GET("/getPost", controllers.PostGetting)
	// r.GET("/getPost/:id", controllers.PostShowByID)
	// r.PUT("/update/:id", controllers.PostUpdate)
	// r.DELETE("/delete/:id", controllers.PostDelete)
	// listen and serve on 0.0.0.0:8080

	api := r.Group("/api")
	{
		post := api.Group("/post")
		{
			post.GET("/", controllers.PostGetting)
			post.GET("/:id", controllers.PostShowByID)
			post.POST("create", controllers.PostCreate)
			post.PUT("/update/:id", controllers.PostUpdate)
			post.DELETE("/delete/:id", controllers.PostDelete)
		}
		author := api.Group("/author")
		{
			author.GET("/", controllers.AuthorGetting)
			author.POST("/create", controllers.AuthorCreate)
		}
		book := api.Group("/book")
		{
			book.GET("/", controllers.BookGetting)
			book.POST("/create", controllers.BookCreate)

		}
	}
	r.Run()
}

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
			author.GET("/:id", controllers.AuthorShowByID)
			author.PUT("/update/:id", controllers.AuthorUpdate)
			author.DELETE("/delete/:id", controllers.AuthorDelete)
		}
		book := api.Group("/book")
		{
			book.GET("/", controllers.BookGetting)
			book.POST("/create", controllers.BookCreate)
			book.GET("/:id", controllers.BookShowByID)
			book.PUT("/update/:id", controllers.BookUpdate)
			book.DELETE("/delete/:id", controllers.BookDelete)
		}
		category := api.Group("/category")
		{
			category.GET("/", controllers.CategoryGetting)
			category.POST("/create", controllers.CategoryCreate)
			category.GET("/:id", controllers.CategoryShowByID)
			category.PUT("/update/:id", controllers.CategoryUpdate)
			category.DELETE("/delete/:id", controllers.CategoryDelete)
		}
		test := api.Group("/books")
		{
			test.GET("/", controllers.BookDetail)
		}

	}
	r.Run()
}

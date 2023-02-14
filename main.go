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

	type x struct {
		Path       string
		Handler    func(*gin.Context)
		HttpMethod string
	}

	// mapx := make(map[string]*x)
	// mapx["book_detail"] = &x{
	// 	Path:       "/bookdetail",
	// 	HttpMethod: "GET",
	// 	Handler:    controllers.BookDetail,
	// }
	// mapx["author_detail"] = &x{
	// 	Path:       "/authordetail",
	// 	HttpMethod: "GET",
	// 	Handler:    controllers.AuthorGetting,
	// }
	// mapx["category_detail"] = &x{
	// 	Path:       "/category_detail",
	// 	HttpMethod: "GET",
	// 	Handler:    controllers.CategoryGetting,
	// }

	var routeConfig = map[string]*x{
		"book_detail": &x{
			Path:       "/book_detail",
			HttpMethod: "GET",
			Handler:    controllers.BookDetail,
		},
		"author_detail": &x{
			Path:       "/author_detail/",
			HttpMethod: "GET",
			Handler:    controllers.AuthorShowByID,
		},
		"category_detail": &x{
			Path:       "/category_detail",
			HttpMethod: "GET",
			Handler:    controllers.CategoryShowByID,
		},
	}

	api := r.Group("/api")
	for _, v := range routeConfig {
		api.Handle(v.HttpMethod, v.Path, v.Handler)
	}
	// for _, v := range mapx {
	// 	api.Handle(v.HttpMethod, v.Path, v.Handler)
	// }
	r.Run()

}

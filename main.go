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
	// type link struct {
	// 	Path    string
	// 	ModelId string
	// }

	routeMap := map[string]string{
		"book":     "book/",
		"author":   "author/",
		"category": "category/",
	}

	var routeConfig = map[string]*x{
		"book_detail": &x{
			Path:       "/book_detail",
			HttpMethod: "GET",
			Handler:    controllers.BookDetail(routeMap),
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
	// var linkConfig = map[string]*link{
	// 	"book_detail": &link{
	// 		Path:    "/book_detail",
	// 		ModelId: strconv.Itoa(1),
	// 	},
	// }

	// var LinkBuilder = func() []x {
	// 	var links []x
	// 	for _, v := range routeConfig {
	// 		links = append(links, *v)
	// 	}
	// 	return links
	// }

	api := r.Group("/api")
	for _, v := range routeConfig {
		api.Handle(v.HttpMethod, v.Path, v.Handler)
	}

	// for _, v := range mapx {
	// 	api.Handle(v.HttpMethod, v.Path, v.Handler)
	// }
	r.Run()

}

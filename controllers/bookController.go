package controllers

import (
	"go-crud/initialize"
	"go-crud/models"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Models []Models

type Link struct {
	Rel  string
	Path string
}

type Meta struct {
	Description string
	Keywords    []string
}

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
		"message": "Book found successfully",
		"error":   false,
		"data":    book,
		"links_related": []gin.H{
			{"method": "GET"},
			{
				"self_URL": "http://localhost:3000/api/book/" + id,
			},
			{
				"authors_URL": "http://localhost:3000/api/author/" + author_ID,
			},
			{
				"categories_URL": "http://localhost:3000/api/category/" + category_ID,
			},
		},
	})
}

func BookShowDetail(c *gin.Context) {
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
		"message": "Book found successfully",
		"error":   false,
		"data":    book,
		"links_related": []gin.H{
			{"method": "GET"},
			{
				"self_URL": "http://localhost:3000/api/book/" + id,
			},
			{
				"authors_URL": "http://localhost:3000/api/author/" + author_ID,
			},
			{
				"categories_URL": "http://localhost:3000/api/category/" + category_ID,
			},
		},
	})
}

func BookBuilder(book *models.Book, rel string) *Link {
	return &Link{
		Rel:  rel,
		Path: "/api/book_detail/" + strconv.Itoa(int(book.ID)),
	}
}

func AuthorBuilder(author *models.Author, rel string) *Link {
	return &Link{
		Rel:  rel,
		Path: "/api/author_detail/" + strconv.Itoa(int(author.ID)),
	}
}

func CategoryBuilder(category *models.Category, rel string) *Link {
	return &Link{
		Rel:  rel,
		Path: "/api/category_detail/" + strconv.Itoa(int(category.ID)),
	}
}

func GetAuthorID(author *models.Author) string {
	author_id := strconv.Itoa(int(models.Book{}.AuthorId))
	return author_id
}

func GetCategoryID(category *models.Category) string {
	category_id := strconv.Itoa(int(models.Book{}.CategoryId))
	return category_id
}

func GetBookID(book *models.Book) string {
	book_id := strconv.Itoa(int(book.ID))
	return book_id
}

var routeMap = map[string]string{
	"book":     "book/",
	"author":   "author/",
	"category": "category/",
}

func getType(in interface{}) string {
	return reflect.TypeOf(in).String()
}

var configMap = map[interface{}]string{
	getType(models.Book{}):     "/api/book_detail/",
	getType(models.Author{}):   "/api/author_detail/",
	getType(models.Category{}): "/api/category_detail/",
}

func buildDetailLink(model interface{}, id string) string {
	prefix := configMap[getType(model)]
	return prefix + id
}

func BookDetail(routeMap map[string]string) func(*gin.Context) {
	return func(c *gin.Context) {
		//id := c.Param("id")
		var books []models.Book
		var authors []models.Author
		var categories []models.Category

		initialize.DB.Find(&books)
		// pp.Fprint(&books)
		initialize.DB.Find(&authors)
		initialize.DB.Find(&categories)
		// initialize.DB.First(&books, id)

		c.IndentedJSON(200, gin.H{
			"data": books,
			"Links": gin.H{
				"_Self": gin.H{
					"method": "GET",
					"self":   buildDetailLink(models.Book{}, GetBookID(&books[9])),
				},
				"Author": gin.H{
					"method": "GET",
					"author": buildDetailLink(models.Author{}, GetAuthorID(&authors[5])),
				},
				"Category": gin.H{
					"method":   "GET",
					"category": buildDetailLink(models.Category{}, GetCategoryID(&categories[1])),
				},
			},
		})
	}
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

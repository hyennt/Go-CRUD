package main

import (
	"go-crud/initialize"
	"go-crud/models"
)

func init() {
	initialize.LoadEnvVar()
	initialize.ConnectDB()
}

func main() {
	initialize.DB.AutoMigrate(&models.Post{})
	initialize.DB.AutoMigrate(&models.Author{})
	initialize.DB.AutoMigrate(&models.Book{})
	initialize.DB.AutoMigrate(&models.Category{})
	initialize.DB.AutoMigrate(&models.User{})
	// drop table
	//initialize.DB.Migrator().DropTable(&models.Post{})
}

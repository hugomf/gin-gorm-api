package main

import (
	"github.com/hugomf/gin-gorm-api/initializers"
	"github.com/hugomf/gin-gorm-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

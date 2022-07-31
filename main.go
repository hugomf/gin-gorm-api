package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hugomf/gin-gorm-api/controllers"
	"github.com/hugomf/gin-gorm-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// router.Use(cors.Default())
	// router.Use(limit.MaxAllowed(200))

	router.POST("/post", controllers.PostCreate)
	router.PUT("/post/:id", controllers.PostUpdate)
	router.GET("/posts", controllers.PostFindAll)
	router.GET("/post/:id", controllers.PostFindById)
	router.DELETE("/post/:id", controllers.PostDelete)

	router.Run() // listen and serve on 0.0.0.0:8080
}

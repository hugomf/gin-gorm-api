package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/hugomf/gin-gorm-api/initializers"
	"github.com/hugomf/gin-gorm-api/models"
)

type Post struct {
	Title string
	Body  string
}

func PostCreate(c *gin.Context) {

	var post Post

	c.BindJSON(&post)

	postModel := models.Post{Title: post.Title, Body: post.Body}
	result := initializers.DB.Create(&postModel)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error,
			"json":    &post,
		})
	}
	jsonData, _ := json.Marshal(postModel)
	c.Data(200, "application/json", jsonData)
}

func PostFindAll(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostFindById(c *gin.Context) {

	var post models.Post

	id := c.Param("id")

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostUpdate(c *gin.Context) {

	var post Post

	c.BindJSON(&post)

	var postModel models.Post

	id := c.Param("id")

	initializers.DB.First(&postModel, id)

	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&postModel).Updates(models.Post{
		Title: post.Title,
		Body:  post.Body,
	})

	c.JSON(200, gin.H{
		"posts": postModel,
	})

}

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)

}

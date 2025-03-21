package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/go-crud/initializers"
	"github.com/scientist-v08/go-crud/models"
)

func PostsCreate(c *gin.Context) {

	// Get data from the data request
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the successful result
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsGetAll(c *gin.Context) {
	// Get all objects from DB
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return all the objects that have been found
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsGetSingle(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	// Get all objects from DB
	var post models.Post
	initializers.DB.First(&post, id)

	// Return all the objects that have been found
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdateAPost(c *gin.Context) {
	// Get the ID from the URl
	id:= c.Param("id")

	// Get all data from the request body
	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	// Get the data already present in the DB
	var post models.Post
	initializers.DB.First(&post, id)

	// Update the data in the DB
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the ID from the URl
	id:= c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/go-crud/initializers"
	"github.com/scientist-v08/go-crud/models"
)

func PostPeople(c *gin.Context) {
	// Get the data from request body
	var reqBody struct {
		FirstName string
		LastName  string
	}
	c.Bind(&reqBody)

	// Save it into database
	personToAdd := models.Person{FirstName: reqBody.FirstName, LastName: reqBody.LastName}
	result := initializers.DB.Create(&personToAdd)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"Failed": "Unable to add a person",
		})
		return
	}

	// Respond the saved data
	c.JSON(200, gin.H{
		"Successfully added the following person": personToAdd,
	})
}

func GetAllPeople(c *gin.Context) {
	// Get all people
	var people []models.Person
	initializers.DB.Find(&people)

	// Return the obtained result
	c.JSON(200, gin.H{
		"people": people,
	})
}

func GetPeoplePaginated(c *gin.Context) {
	// Get query parameters
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))

	// Calculate offset
	offset := (pageNumber - 1) * pageSize

	// Get all people with pagination
	var people []models.Person
	initializers.DB.Limit(pageSize).Offset(offset).Find(&people)

	// Get the total number of elements
	var count int64
	initializers.DB.Model(&models.Person{}).Count(&count)

	// Return the obtained result
	c.JSON(200, gin.H{
		"people": people,
		"totalElements": count,
	})
}
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/go-crud/initializers"
	"github.com/scientist-v08/go-crud/models"
)

func AddProduct(c *gin.Context) {
	// Get the data from the request body
	var reqBody struct {
		Code string
		Name string
		Category string
		Quantity string
	}
	c.Bind(&reqBody)

	// Create a product variable to save in the DB
	productToAdd := models.Products{
		Code: reqBody.Code,
		Name: reqBody.Name,
		Category: reqBody.Category,
		Quantity: reqBody.Quantity,
	}

	// Now save the data
	result := initializers.DB.Create(&productToAdd)

	// Check if there was an error when we tried to save the data
	if result.Error != nil {
		c.JSON(400, gin.H{
			"Failed": "Could not save data",
		})
		return 
	}

	// If there was no error return the success message
	c.JSON(200, gin.H{
		"Successfully added a product": productToAdd,
	})
}

func GetAllProducts(c *gin.Context) {
	// Create a local variable that stores all the data we are going to obtain from the DB
	var allProducts []models.Products
	initializers.DB.Find(&allProducts)

	c.JSON(200, gin.H{
		"products": allProducts,
	})
}
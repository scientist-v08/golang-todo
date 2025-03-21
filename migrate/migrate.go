package main

import (
	"github.com/scientist-v08/go-crud/initializers"
	"github.com/scientist-v08/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
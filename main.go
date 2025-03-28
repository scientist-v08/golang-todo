package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/go-crud/controllers"
	"github.com/scientist-v08/go-crud/initializers"
	"github.com/scientist-v08/go-crud/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200", "http://localhost:4201"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/create/user", controllers.SignUp)
	r.POST("/login/user", controllers.Login)
	r.POST("/posts", middleware.RequireAnyRole("ROLE_USER"), controllers.PostsCreate)
	r.PUT("/posts/:id", middleware.RequireAnyRole("ROLE_ADMIN"), controllers.PostsUpdateAPost)
	r.GET("/posts", middleware.RequireAnyRole("ROLE_USER"), controllers.PostsGetAll)
	r.GET("/posts/:id", middleware.RequireAnyRole("ROLE_USER"), controllers.PostsGetSingle)
	r.DELETE("/posts/:id", middleware.RequireAnyRole("ROLE_ADMIN"), controllers.PostsDelete)
	r.POST("/person", middleware.RequireAnyRole("ROLE_USER"), controllers.PostPeople)
	r.GET("/get/people", middleware.RequireAnyRole("ROLE_USER"), controllers.GetAllPeople)
	r.GET("/get/people/paginated", middleware.RequireAnyRole("ROLE_USER"), controllers.GetPeoplePaginated)
	r.POST("/products/create", controllers.AddProduct)
	r.GET("/products", controllers.GetAllProducts)
	r.Run()
}
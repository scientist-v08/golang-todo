package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/scientist-v08/go-crud/initializers"
	"github.com/scientist-v08/go-crud/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get the email or password from the request body
	var user struct {
		Email string
		Password string
	}
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read new body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the User
	newUser := models.User{Email: user.Email, Password: string(hash), Roles: []string{"ROLE_USER"}}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a new user",
		})
		return
	}

	// Respond
	c.JSON(200, gin.H{
		"Creating a new user": "Successful",
	})
}

func Login(c *gin.Context) {
	// Get the email and password from the request body
	var user struct {
		Email string
		Password string
	}
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Look up requested user
	var existingUser models.User
	initializers.DB.First(&existingUser, "email = ?", user.Email)

	if existingUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email ID",
		})
		return
	}

	// Compare password from request body to saved password in DB
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))

	if bcryptErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Generate a JWT token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := os.Getenv("JWT_SECRET")
	tokenString, jwtTokenErr := token.SignedString([]byte(hmacSampleSecret))

	if jwtTokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create JWT token",
		})
		return
	}

	// Send it back
	c.JSON(200, gin.H{
		"access_token": tokenString,
	})
}
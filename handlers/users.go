package handlers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/database"
	"github.com/hankrugg/CollegeCooks/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ListUsers(c *gin.Context) {
	// Define a slice to hold the retrieved facts
	var users []models.User

	// Retrieve all facts from the database
	if err := database.DB.Db.Find(&users).Error; err != nil {
		// If an error occurs, return an internal server error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Return the retrieved facts as a response
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	// Define a struct to represent the JSON body
	var requestBody struct {
		Email     string `json:"email"`
		FirstName string `json:"first"`
		LastName  string `json:"last"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	// Bind the JSON body to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Now you can access the values from the JSON body
	Email := requestBody.Email
	FirstName := requestBody.FirstName
	LastName := requestBody.LastName
	Username := requestBody.Username
	Password, _ := hashPassword(requestBody.Password)

	// Use the values as needed

	// For example, you might create a new Fact with these values:
	user := models.User{Email: Email, FirstName: FirstName, LastName: LastName, Username: Username, Password: Password}
	tx := database.DB.Db.Create(&user)

	fmt.Println(tx)

	// Return a response
	c.JSON(http.StatusOK, gin.H{
		"message":  "User created successfully",
		"email":    Email,
		"username": Username,
		"password": Password,
	})
}

func hashPassword(rawPassword string) (string, error) {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Encode the hashed password bytes using base64
	encodedHash := base64.StdEncoding.EncodeToString(hashedPassword)
	return encodedHash, nil
}

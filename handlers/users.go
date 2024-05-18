package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hankrugg/CollegeCooks/database"
	"github.com/hankrugg/CollegeCooks/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
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

func Register(c *gin.Context) {
	// Define a struct to represent the JSON body
	var requestBody struct {
		Email     string `json:"email"`
		FirstName string `json:"first"`
		LastName  string `json:"last"`
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
	Password, _ := hashPassword(requestBody.Password)

	user := models.User{Email: Email, FirstName: FirstName, LastName: LastName, Password: string(Password)}
	err := database.DB.Db.Create(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating user",
			"email":   Email,
		})
		return
	}

	// Return a response
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"email":   Email,
	})
	return
}

func hashPassword(rawPassword string) ([]byte, error) {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return []byte(""), err
	}
	// Encode the hashed password bytes using base64
	return hashedPassword, nil
}

func Login(c *gin.Context) {
	// find user
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&requestBody)
	if err != nil {
		// Handle error
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to read request data",
		})
		return
	}

	var user models.User

	database.DB.Db.First(&user, "email = ?", requestBody.Email)
	if user.ID == 0 {
		// If an error occurs, return an internal server error response
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User and password combination not found",
		})
		return
	}

	//validate user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User and password combination not found",
		})
		return
	}

	// create jwt
	expiration := time.Now().Add(time.Hour)
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiration),
		Issuer:    "test",
		Subject:   strconv.Itoa(int(user.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingKey := []byte(os.Getenv("SIGNING_KEY")) // Load the signing key from environment variable
	ss, err := token.SignedString(signingKey)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to grant token",
		})
		return
	}

	// create cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", ss, expiration.Minute(), "", "", false, true)
	// respond

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
	})

}

func Validate(c *gin.Context) {
	
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

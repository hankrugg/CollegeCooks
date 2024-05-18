package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hankrugg/CollegeCooks/database"
	"github.com/hankrugg/CollegeCooks/models"
	"net/http"
	"os"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	database.ConnectDb()

	setupRoutes(r)

	r.Run(":3000")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func RequireAuth(c *gin.Context) {
	ss, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(ss, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token is valid
		expFloat, ok := claims["exp"].(float64)
		if !ok {
			// Handle invalid expiration time format
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		expUnix := int64(expFloat)

		// Compare expiration time with current Unix time
		if expUnix <= time.Now().Unix() {
			// Token is expired
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user
		var user models.User
		if err := database.DB.Db.First(&user, claims["sub"]).Error; err != nil {
			// User not found or database error
			// Log the error or handle it appropriately
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set the user information in the request context
		c.Set("user", user)

		// Continue processing the request
		c.Next()
	} else {
		// Token is invalid or claims type assertion failed
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

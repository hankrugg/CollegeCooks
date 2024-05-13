package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/database"
	"net/http"
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

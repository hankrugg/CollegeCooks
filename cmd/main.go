package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/database"
)

func main() {
	r := gin.Default()

	database.ConnectDb()

	setupRoutes(r)

	r.Run(":3000")
}

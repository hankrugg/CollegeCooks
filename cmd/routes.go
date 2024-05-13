package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/handlers"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/", handlers.ListUsers)

	r.POST("/addUser", handlers.CreateUser)

}

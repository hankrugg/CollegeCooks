package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/handlers"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/getUsers", handlers.ListUsers)

	r.POST("/addUser", handlers.CreateUser)

}

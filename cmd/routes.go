package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/handlers"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/getUser", RequireAuth, handlers.GetUser)

	r.POST("/register", handlers.Register)

	r.POST("/login", handlers.Login)

	//r.GET("/validate", RequireAuth, handlers.Validate)

	r.POST("/addIngredient", RequireAuth, handlers.AddIngredient)

	r.GET("/getIngredients", RequireAuth, handlers.GetIngredients)

	r.POST("/addRecipe", RequireAuth, handlers.AddRecipe)

	r.GET("/getRecipes", RequireAuth, handlers.GetRecipes)

}

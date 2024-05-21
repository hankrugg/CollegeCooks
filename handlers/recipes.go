package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/database"
	"github.com/hankrugg/CollegeCooks/models"
	"net/http"
)

func AddRecipe(c *gin.Context) {

	var requestBody struct {
		Title       string `json:"title"`
		Directions  string `json:"directions"`
		Time        string `json:"time"`
		Ingredients string `json:"ingredients"`
		Source      string `json:"source"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Now you can access the values from the JSON body
	Title := requestBody.Title
	Directions := requestBody.Directions
	Time := requestBody.Time
	Ingredients := requestBody.Ingredients
	Source := requestBody.Source

	recipe := models.Recipe{Title: Title, Directions: Directions, Time: Time, Ingredients: Ingredients, Source: Source}
	err := database.DB.Db.Create(&recipe).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating ingredient",
			"name":    Title,
		})
		return
	}

	// Return a response
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe added successfully",
		"name":    Title,
	})
	return
}

func GetRecipes(c *gin.Context) {

	// Retrieve all facts from the database
	var recipes []models.Recipe
	err := database.DB.Db.Find(&recipes).Error
	if err != nil {
		// Handle the error, such as returning a suitable response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find ingredients"})
		return
	}

	// Return the retrieved facts as a response
	c.JSON(http.StatusOK, gin.H{
		"recipes": recipes,
	})
}

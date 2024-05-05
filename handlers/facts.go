package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/hankrugg/CollegeCooks/database"
	"github.com/hankrugg/CollegeCooks/models"
	"net/http"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Div Rhino Trivia App!")
}

func ListFacts(c *gin.Context) {
	// Define a slice to hold the retrieved facts
	var facts []models.Fact

	// Retrieve all facts from the database
	if err := database.DB.Db.Find(&facts).Error; err != nil {
		// If an error occurs, return an internal server error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Return the retrieved facts as a response
	c.JSON(http.StatusOK, facts)
}

func CreateFact(c *gin.Context) {
	// Define a struct to represent the JSON body
	var requestBody struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	// Bind the JSON body to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Now you can access the values from the JSON body
	question := requestBody.Question
	answer := requestBody.Answer

	// Use the values as needed

	// For example, you might create a new Fact with these values:
	fact := models.Fact{Question: question, Answer: answer}
	database.DB.Db.Create(&fact)

	// Return a response
	c.JSON(http.StatusOK, gin.H{
		"message":  "Fact created successfully",
		"question": question,
		"answer":   answer,
	})
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hankrugg/CollegeCooks/database"
	"github.com/hankrugg/CollegeCooks/models"
	"net/http"
)

func AddIngredient(c *gin.Context) {

	var requestBody struct {
		Name      string      `json:"Name" gorm:"text;not null;default:null"`
		Quantity  int         `json:"Quantity" gorm:"int;not null;default:1"`
		Fruit     bool        `json:"fruit" gorm:"bool; not null;default:false"`
		Vegetable bool        `json:"vegetable" gorm:"bool; not null;default:false"`
		Meat      bool        `json:"meat" gorm:"bool; not null;default:false"`
		Grain     bool        `json:"grain" gorm:"bool; not null;default:false"`
		UserID    uint        `json:"UserID" gorm:"not null"` // Foreign key
		User      models.User `json:"user"`                   // Associated user
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Now you can access the values from the JSON body
	Name := requestBody.Name
	Quantity := requestBody.Quantity
	Fruit := requestBody.Fruit
	Vegetable := requestBody.Vegetable
	Meat := requestBody.Meat
	Grain := requestBody.Grain
	UserID := requestBody.UserID
	User := requestBody.User

	ingredient := models.Ingredient{Name: Name, Quantity: Quantity, Fruit: Fruit, Vegetable: Vegetable, Meat: Meat, Grain: Grain, UserID: UserID, User: User}
	err := database.DB.Db.Create(&ingredient).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating ingredient",
			"name":    Name,
		})
		return
	}

	// Return a response
	c.JSON(http.StatusOK, gin.H{
		"message": "Ingredient added successfully",
		"name":    Name,
	})
	return
}

func GetIngredients(c *gin.Context) {

	user, _ := c.Get("user")
	userID := user.(models.User).ID

	// Retrieve all facts from the database
	var ingredients []models.Ingredient
	err := database.DB.Db.Where("user_id = ?", userID).Find(&ingredients).Error
	if err != nil {
		// Handle the error, such as returning a suitable response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find ingredients"})
		return
	}

	// Return the retrieved facts as a response
	c.JSON(http.StatusOK, gin.H{
		"ingredients": ingredients,
	})
}

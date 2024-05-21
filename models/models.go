package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"text;not null;default:null;unique"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Password  string `json:"password"`
}

type Recipe struct {
	gorm.Model
	Title       string `json:"title" gorm:"text;not null;default:null;unique"`
	Directions  string `json:"directions"`
	Time        string `json:"time"`
	Ingredients string `json:"ingredients"`
	Source      string `json:"source"`
}

type Ingredient struct {
	gorm.Model
	Name      string `json:"Name" gorm:"text;not null;default:null"`
	Quantity  int    `json:"Quantity" gorm:"int;not null;default:1"`
	Fruit     bool   `json:"fruit" gorm:"bool; not null;default:false"`
	Vegetable bool   `json:"vegetable" gorm:"bool; not null;default:false"`
	Meat      bool   `json:"meat" gorm:"bool; not null;default:false"`
	Grain     bool   `json:"grain" gorm:"bool; not null;default:false"`
	UserID    uint   `json:"UserID" gorm:"not null"` // Foreign key
	User      User   `json:"user"`                   // Associated user
}

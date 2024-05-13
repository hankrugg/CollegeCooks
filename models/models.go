package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Recipe struct {
	gorm.Model
	Ingredients []Ingredient `json:"Ingredients"`
	Directions  string       `json:"Directions"`
	PrepTime    int          `json:"PrepTime"`
	CookTime    int          `json:"CookTime"`
}

type Ingredient struct {
	gorm.Model
	Name     string `json:"Name" gorm:"text;not null;default:null`
	Fruit    bool   `json:"fruit" gorm:"bool; not null;default:null`
	LastName string `json:"last" gorm:"text; not null;default:null`
	Username string `json:"username" gorm:"text;not null;default:null`
}

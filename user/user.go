package user

import "github.com/jinzhu/gorm"

// User is a Human
type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

// Users is a slice for User
// var Users []User

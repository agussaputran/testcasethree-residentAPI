package models

import "gorm.io/gorm"

// Users Model
type Users struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

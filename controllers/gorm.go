package controllers

import "gorm.io/gorm"

// Gorm controller support
type Gorm struct {
	DB *gorm.DB
}

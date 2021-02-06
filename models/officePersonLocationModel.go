package models

import "gorm.io/gorm"

// OfficePersonLocations model
type OfficePersonLocations struct {
	gorm.Model
	OfficeID uint
	PersonID uint
}

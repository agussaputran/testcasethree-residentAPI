package models

import "gorm.io/gorm"

// Offices model
type Offices struct {
	gorm.Model
	Name                 string
	SubDistrictID        uint
	OfficePersonLocation []OfficePersonLocations `gorm:"ForeignKey:OfficeID"`
}

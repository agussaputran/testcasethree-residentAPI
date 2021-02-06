package models

import "gorm.io/gorm"

// Persons model
type Persons struct {
	gorm.Model
	Nip, FullName, FirstName, LastName, BirthDate, BirthPlace, PhotoProfileUrl, Gender, ZoneLocation string
	SubDistrictID                                                                                    uint
	OfficePersonLocation                                                                             []OfficePersonLocations `gorm:"ForeignKey:PersonID"`
}

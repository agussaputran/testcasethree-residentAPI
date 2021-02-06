package seeders

import (
	"fmt"
	"strconv"
	"testcasethree-residentAPI/models"

	"gorm.io/gorm"
)

// SeedOfficePersonLocation func
func SeedOfficePersonLocation(db *gorm.DB) {
	var ofcPersonArr = [...][2]string{
		{"1", "3"},
		{"1", "9"},
		{"4", "1"},
		{"6", "12"},
		{"5", "2"},
		{"3", "8"},
		{"1", "8"},
		{"6", "2"},
		{"3", "9"},
		{"3", "5"},
		{"3", "6"},
		{"3", "7"},
		{"2", "12"},
		{"6", "13"},
		{"5", "14"},
		{"6", "22"},
		{"2", "21"},
		{"4", "18"},
		{"1", "14"},
		{"3", "14"},
	}

	var officePerson models.OfficePersonLocations
	for _, v1 := range ofcPersonArr {
		office, _ := strconv.ParseUint(v1[0], 10, 32)
		officePerson.OfficeID = uint(office)
		person, _ := strconv.ParseUint(v1[1], 10, 32)
		officePerson.PersonID = uint(person)
		officePerson.ID = 0
		db.Create(&officePerson)

	}
	fmt.Println("Seeder OfficePersonLocation created")
}

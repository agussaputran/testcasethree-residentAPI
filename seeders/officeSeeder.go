package seeders

import (
	"fmt"
	"strconv"
	"testcasethree-residentAPI/models"

	"gorm.io/gorm"
)

// SeedOffice func
func SeedOffice(db *gorm.DB) {
	var officeArray = [...][2]string{
		{"8", "Kantor Camat Jatiasih"},
		{"1", "PT Privy Identitas Digital"},
		{"7", "Panasonic Group Indonesia"},
		{"1", "PT Xapiens Teknologi Indonesia"},
		{"2", "Javaloka IT Consultant"},
		{"1", "JMC IT Consultant"},
	}

	var office models.Offices
	for _, v1 := range officeArray {
		data, _ := strconv.ParseUint(v1[0], 10, 32)
		office.SubDistrictID = uint(data)
		office.Name = v1[1]
		office.ID = 0
		db.Create(&office)

	}
	fmt.Println("Seeder Office created")
}

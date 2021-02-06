package seeders

import (
	"fmt"
	"strconv"
	"testcasethree-residentAPI/models"

	"gorm.io/gorm"
)

// SeedSubDistrict func
func SeedSubDistrict(db *gorm.DB) {
	var subDistrictArray = [...][2]string{
		{"39", "Bantul"},
		{"39", "Banguntapan"},
		{"39", "Imogiri"},
		{"55", "Bekasi Barat"},
		{"55", "Bekasi Selatan"},
		{"55", "Bekasi Timur"},
		{"55", "Bekasi Utara"},
		{"55", "Jatiasih"},
		{"55", "Jatisampurna"},
		{"55", "Pondok Gede"},
		{"55", "Pondok Melati"},
		{"55", "Rawalumbu"},
		{"79", "Bogor Barat"},
		{"79", "Bogor Selatan"},
		{"79", "Bogor Tengah"},
		{"79", "Bogor Timur"},
		{"79", "Bogor Utara"},
	}

	var subDist models.SubDistricts
	for _, v1 := range subDistrictArray {
		data, _ := strconv.ParseUint(v1[0], 10, 32)
		subDist.DistrictID = uint(data)
		subDist.Name = v1[1]
		subDist.ID = 0
		db.Create(&subDist)

	}
	fmt.Println("Seeder Sub District created")
}

package seeders

import (
	"fmt"
	"log"
	"testcasethree-residentAPI/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUser func
func SeedUser(db *gorm.DB) {
	var userArray = [...][3]string{
		{"admin@xapiens.id", "admin", "admin"},
		{"data@xapiens.id", "data", "entry"},
		{"agussaputra@gmail.com", "agussaputra", "guest"},
	}

	var user models.Users
	for _, v1 := range userArray {
		user.Email = v1[0]
		user.Password = v1[1]
		user.Role = v1[2]
		user.ID = 0

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error -> ", err.Error())
		}
		user.Password = string(hash)

		db.Create(&user)

	}
	fmt.Println("Seeder User created")
}

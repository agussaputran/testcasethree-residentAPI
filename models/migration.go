package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var (
		checkProvinces    bool
		checkDistricts    bool
		checkSubDistricts bool
		checkPersons      bool
		checkUsers        bool
	)

	db.Migrator().DropTable(&Provinces{})
	db.Migrator().DropTable(&Districts{})
	db.Migrator().DropTable(&SubDistricts{})
	db.Migrator().DropTable(&Persons{})
	db.Migrator().DropTable(&Users{})

	checkProvinces = db.Migrator().HasTable(&Provinces{})
	if !checkProvinces {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("Create Provinces Table")
	}

	checkDistricts = db.Migrator().HasTable(&Districts{})
	if !checkDistricts {
		db.Migrator().CreateTable(&Districts{})
		fmt.Println("Create Districts Table")
	}

	checkSubDistricts = db.Migrator().HasTable(&SubDistricts{})
	if !checkSubDistricts {
		db.Migrator().CreateTable(&SubDistricts{})
		fmt.Println("Create SubDistricts Table")
	}

	checkPersons = db.Migrator().HasTable(&Persons{})
	if !checkPersons {
		db.Migrator().CreateTable(&Persons{})
		fmt.Println("Create Persons Table")
	}

	checkUsers = db.Migrator().HasTable(&Users{})
	if !checkUsers {
		db.Migrator().CreateTable(&Users{})
		fmt.Println("Create Users Table")
	}
}

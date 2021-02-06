package seeders

import (
	"fmt"
	"strconv"
	"testcasethree-residentAPI/models"

	"gorm.io/gorm"
)

// SeedPerson func
func SeedPerson(db *gorm.DB) {
	var personArray = [...][10]string{
		{"1", "123", "Agus Saputra", "Agus", "Saputra", "1997-08-08", "Gorontalo", "http://someplace.com/agus.jpg", "M", "WIB"},
		{"2", "321", "Gabriela Sin", "Gabriela", "Sin", "1995-05-25", "Yogya", "http://someplace.com/gabriela.jpg", "F", "WIB"},
		{"8", "567", "Adit Rahmat", "Adit", "Rahmat", "1997-02-14", "Bekasi", "http://someplace.com/adit.jpg", "M", "WIB"},
		{"3", "223", "Indah Pramesti", "Indah", "Pramesti", "1992-03-15", "Bantul", "http://someplace.com/indah.jpg", "F", "WIB"},
		{"4", "352", "Derbi Romero", "Derbi", "Romero", "1993-04-12", "Bekasi", "http://someplace.com/derbi.jpg", "M", "WIB"},
		{"5", "212", "Juki Suherman", "Juki", "Suherman", "1994-05-13", "Bekasi", "http://someplace.com/juki.jpg", "M", "WIB"},
		{"6", "732", "Fella Andini", "Fella", "Andini", "1995-07-18", "Bekasi", "http://someplace.com/fella.jpg", "F", "WIB"},
		{"7", "978", "Anindita Licia", "Anindita", "Licia", "1998-03-12", "Bekasi", "http://someplace.com/anin.jpg", "F", "WIB"},
		{"8", "421", "Juminten", "Juminten", "", "1990-02-10", "Bekasi", "http://someplace.com/juminten.jpg", "F", "WIB"},
		{"9", "754", "Irfan Sampurna", "Irfan", "Sampurna", "1988-04-12", "Bekasi", "http://someplace.com/irfan.jpg", "M", "WIB"},
		{"15", "827", "Santi Resa", "Santi", "Resa", "1987-02-28", "Bogor", "http://someplace.com/santi.jpg", "F", "WIB"},
		{"2", "632", "Farhan Ibrahim", "Farhan", "Ibrahim", "1992-03-08", "Bantul", "http://someplace.com/farhan.jpg", "M", "WIB"},
		{"1", "740", "Josh Dun", "Josh", "Dun", "1997-03-08", "Bantul", "http://someplace.com/josh.jpg", "M", "WIB"},
		{"8", "937", "Arman Iskandar", "Arman", "Iskandar", "1987-03-02", "Bekasi", "http://someplace.com/arman.jpg", "M", "WIB"},
		{"1", "087", "Faruk Iskandar", "Faruk", "Iskandar", "1997-03-08", "Bantul", "http://someplace.com/faruk.jpg", "M", "WIB"},
		{"7", "876", "Herman Adam", "Herman", "Adam", "1996-02-07", "Bantul", "http://someplace.com/herman.jpg", "M", "WIB"},
		{"7", "653", "Geby Sinatra", "Geby", "Sinatra", "1998-02-02", "Bekasi", "http://someplace.com/faruk.jpg", "F", "WIB"},
		{"8", "823", "Sinta Putri", "Sinta", "Putri", "1997-04-04", "Bantul", "http://someplace.com/sinta.jpg", "F", "WIB"},
		{"2", "987", "Gery Fren", "Gery", "Fren", "1993-11-23", "Bantul", "http://someplace.com/gery.jpg", "M", "WIB"},
		{"1", "836", "Abdul Karim", "Abdul", "Karim", "1996-06-06", "Bantul", "http://someplace.com/abdul.jpg", "M", "WIB"},
		{"2", "429", "Dandi Firhan", "Dandi", "Firhan", "1992-10-12", "Bantul", "http://someplace.com/dandi.jpg", "M", "WIB"},
		{"1", "938", "Dwi Putri", "Dwi", "Putri", "1991-11-21", "Bantul", "http://someplace.com/dwi.jpg", "F", "WIB"},
	}

	var person models.Persons
	for _, v1 := range personArray {
		data, _ := strconv.ParseUint(v1[0], 10, 32)
		person.SubDistrictID = uint(data)
		person.Nip = v1[1]
		person.FullName = v1[2]
		person.FirstName = v1[3]
		person.LastName = v1[4]
		person.BirthDate = v1[5]
		person.BirthPlace = v1[6]
		person.PhotoProfileUrl = v1[7]
		person.Gender = v1[8]
		person.ZoneLocation = v1[9]
		person.ID = 0
		db.Create(&person)

	}
	fmt.Println("Seeder person created")
}

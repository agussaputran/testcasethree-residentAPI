package controllers

import (
	"fmt"
	"net/http"
	"testcasethree-residentAPI/helper"
	"testcasethree-residentAPI/models"

	"github.com/gin-gonic/gin"
)

type personResponse struct {
	ID uint
	FullName, FirstName, LastName, BirthDate,
	BirthPlace, PhotoProfileUrl, Gender, ZoneLocation, Subdistrict,
	District, Province string
}

// PostCreatePerson route struct method
func (gorm *Gorm) PostCreatePerson(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	if err := c.Bind(&person); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Create(&person)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":              person.ID,
				"Nip":             person.Nip,
				"fullName":        person.FullName,
				"firstName":       person.FirstName,
				"lastName":        person.LastName,
				"subDistrictID":   person.SubDistrictID,
				"birthDate":       person.BirthDate,
				"birthPlace":      person.BirthPlace,
				"photoProfileUrl": person.PhotoProfileUrl,
				"gender":          person.Gender,
				"zoneLocation":    person.ZoneLocation,
				"created_at":      person.CreatedAt,
				"update_at":       person.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadPerson route func
func (gorm *Gorm) GetReadPerson(c *gin.Context) {
	var (
		person   []models.Persons
		response []personResponse
		result   gin.H
	)

	gorm.DB.Model(&person).Select(`persons.id, persons.full_name, persons.first_name,
	persons.last_name, persons.birth_date, persons.birth_place, person.photo_profile_url,
	persons.gender, persons.zone_location, sub_districts.name as subdistrict,
	districts.name as district, provinces.name as province`).Joins(`left join sub_districts
	on sub_districts.id = persons.sub_district_id left join districts on districts.id =
	sub_districts.district_id left join provinces on provinces.id = districts.province_id`).Scan(&response)
	if length := len(response); length <= 0 {
		result = helper.ResultAPINilResponse(response, length)
	} else {
		result = helper.ResultAPIResponse(response, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdatePerson route struct method
func (gorm *Gorm) PatchUpdatePerson(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&person); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Model(&person).Where("id=?", id).Updates(models.Persons{FullName: person.FullName, FirstName: person.FirstName, LastName: person.LastName, PhotoProfileUrl: person.PhotoProfileUrl, SubDistrictID: person.SubDistrictID, ZoneLocation: person.ZoneLocation})
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"fullName":        person.FullName,
				"firstName":       person.FirstName,
				"lastName":        person.LastName,
				"subDistrictID":   person.SubDistrictID,
				"birthDate":       person.BirthDate,
				"birthPlace":      person.BirthPlace,
				"photoProfileUrl": person.PhotoProfileUrl,
				"gender":          person.Gender,
				"zoneLocation":    person.ZoneLocation,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemovePerson route struct method
func (gorm *Gorm) DeleteRemovePerson(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	id := c.Query("id")
	gorm.DB.Delete(&person, id)
	result = gin.H{
		"Message": "Success delete person",
	}
	c.JSON(http.StatusOK, result)
}

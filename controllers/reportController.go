package controllers

import (
	"net/http"
	"testcasethree-residentAPI/helper"
	"testcasethree-residentAPI/models"

	"github.com/gin-gonic/gin"
)

// ReportPersonByGender func route
func (gorm *Gorm) ReportPersonByGender(c *gin.Context) {
	var (
		person models.Persons
		count  int64
		countM int64
		countF int64
		result gin.H
	)

	totalPerson := c.Query("person")

	if totalPerson == "" {
		result = gin.H{
			"message": "should input query",
		}
		c.JSON(http.StatusBadRequest, result)
	} else if totalPerson == "M" || totalPerson == "m" {
		gorm.DB.Model(&person).Where("gender = ?", "M").Count(&count)
		result = gin.H{
			"message":           "success",
			"total_male_person": count,
		}
		c.JSON(http.StatusOK, result)
	} else if totalPerson == "F" || totalPerson == "f" {
		gorm.DB.Model(&person).Where("gender = ?", "F").Count(&count)
		result = gin.H{
			"message":             "success",
			"total_female_person": count,
		}
		c.JSON(http.StatusOK, result)
	} else if totalPerson == "all" {
		gorm.DB.Model(&person).Where("gender = ?", "M").Count(&countM)
		gorm.DB.Model(&person).Where("gender = ?", "F").Count(&countF)
		gorm.DB.Model(&person).Count(&count)
		result = gin.H{
			"message":             "success",
			"total_person":        count,
			"total_male_person":   countM,
			"total_female_person": countF,
		}
		c.JSON(http.StatusOK, result)
	} else {
		result = gin.H{
			"message": "not a valid value",
		}
		c.JSON(http.StatusBadRequest, result)
	}

}

// ReportPersonOffice func route
func (gorm *Gorm) ReportPersonOffice(c *gin.Context) {
	type response struct {
		ID       uint     `json:"id"`
		FullName string   `json:"full_name"`
		Total    int64    `json:"total"`
		CityName []string `json:"city_name"`
	}

	var (
		// person []models.Persons
		result gin.H
		resp   []response
		// resp2  res2
		id                 uint
		FullName, CityName string
		total              int64
	)

	rows, _ := gorm.DB.Table("persons").Select("persons.id as id, persons.full_name as FullName, count(*) as Total").
		Joins(`join office_person_locations opl on opl.person_id = persons.id
	join offices ofc on opl.office_id = ofc.id
	join sub_districts sd on sd.id = ofc.sub_district_id
	join districts d on d.id = sd.district_id`).Group("persons.id, persons.full_name").Rows()
	for rows.Next() {
		rows.Scan(&id, &FullName, &total)
		// fmt.Println(id, " | ", FullName, " | ", total)
		var location []string
		innerRows, _ := gorm.DB.Table("persons").Select("d.name as CityName").Where("persons.id = ?", id).
			Joins(`join office_person_locations opl on opl.person_id = persons.id
		join offices ofc on opl.office_id = ofc.id
		join sub_districts sd on sd.id = ofc.sub_district_id
		join districts d on d.id = sd.district_id`).Rows()

		for innerRows.Next() {
			innerRows.Scan(&CityName)
			// fmt.Println(CityName)
			location = append(location, CityName)
		}
		resp = append(resp, response{id, FullName, total, location})
	}

	// gorm.DB.Preload("OfficePersonLocation").Find(&person)
	if length := len(resp); length <= 0 {
		result = helper.ResultAPINilResponse(resp, length)
	} else {
		result = helper.ResultAPIResponse(resp, length)
	}

	// var data []response
	// result = gin.H{
	// 	"data": resp,
	// }

	c.JSON(http.StatusOK, result)
}

// ReportCountPersonOfficeByGender func route
func (gorm *Gorm) ReportCountPersonOfficeByGender(c *gin.Context) {
	type response struct {
		OfficeName      string
		CountMalePerson int64
	}

	var (
		resp   []response
		office []models.Offices
		result gin.H
	)

	gorm.DB.Model(&office).Where("persons.gender = ?", "M").Select("offices.name as office_name, COUNT(persons.gender) as count_male_person").
		Joins("JOIN office_person_locations ON offices.id = office_person_locations.office_id JOIN persons ON office_person_locations.person_id = persons.id").
		Group("offices.name").
		Scan(&resp)

	// gorm.DB.Preload("OfficePersonLocation").Find(&resp)
	if length := len(resp); length <= 0 {
		result = helper.ResultAPINilResponse(resp, length)
	} else {
		result = helper.ResultAPIResponse(resp, length)
	}

	c.JSON(http.StatusOK, result)
}

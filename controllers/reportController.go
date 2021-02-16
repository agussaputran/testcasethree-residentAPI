package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"testcasethree-residentAPI/helper"
	"testcasethree-residentAPI/models"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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
		result             gin.H
		resp               []response
		id                 uint
		FullName, CityName string
		total              int64
	)

	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10)
	pool.MaxActive = 10

	conn := pool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", "report_response"))
	if err != nil {
		rows, _ := gorm.DB.Table("persons").Select("persons.id as id, persons.full_name as FullName, count(*) as Total").
			Joins(`join office_person_locations opl on opl.person_id = persons.id
		join offices ofc on opl.office_id = ofc.id
		join sub_districts sd on sd.id = ofc.sub_district_id
		join districts d on d.id = sd.district_id`).Group("persons.id, persons.full_name").Rows()

		for rows.Next() {
			rows.Scan(&id, &FullName, &total)
			var location []string
			innerRows, _ := gorm.DB.Table("persons").Select("d.name as CityName").Where("persons.id = ?", id).
				Joins(`join office_person_locations opl on opl.person_id = persons.id
			join offices ofc on opl.office_id = ofc.id
			join sub_districts sd on sd.id = ofc.sub_district_id
			join districts d on d.id = sd.district_id`).Rows()

			for innerRows.Next() {
				innerRows.Scan(&CityName)
				location = append(location, CityName)
			}
			resp = append(resp, response{id, FullName, total, location})
		}
		jd, _ := json.Marshal(resp)
		_, _ = conn.Do("SET", "report_response", jd)
	} else if reply != nil {
		err := json.Unmarshal(reply, &resp)
		if err != nil {
			log.Fatalln("error ->", err.Error())
		}
	}

	_, _ = conn.Do("EXPIRE", "report_response", "120")

	if length := len(resp); length <= 0 {
		result = helper.ResultAPINilResponse(resp, length)
	} else {
		result = helper.ResultAPIResponse(resp, length)
	}
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

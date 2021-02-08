package controllers

import (
	"net/http"
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

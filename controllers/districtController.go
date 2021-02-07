package controllers

import (
	"fmt"
	"net/http"
	"testcasethree-residentAPI/helper"
	"testcasethree-residentAPI/models"

	"github.com/gin-gonic/gin"
)

type districtResponse struct {
	ID       uint
	District string
	Province string
}

// PostCreateDistrict route struct method
func (gorm *Gorm) PostCreateDistrict(c *gin.Context) {
	var (
		district models.Districts
		result   gin.H
	)

	if err := c.Bind(&district); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Create(&district)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":          district.ID,
				"province_id": district.ProvinceID,
				"district":    district.Name,
				"created_at":  district.CreatedAt,
				"update_at":   district.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadDistrict route func
func (gorm *Gorm) GetReadDistrict(c *gin.Context) {
	var (
		district []models.Districts
		response []districtResponse
		result   gin.H
	)

	gorm.DB.Model(&district).Select("districts.id, districts.name as district, provinces.name as province").Joins("left join provinces on provinces.id = districts.province_id").Scan(&response)
	if length := len(response); length <= 0 {
		result = helper.ResultAPINilResponse(response, length)
	} else {
		result = helper.ResultAPIResponse(response, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdateDistrict route struct method
func (gorm *Gorm) PatchUpdateDistrict(c *gin.Context) {
	var (
		district models.Districts
		result   gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&district); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Model(&district).Where("id = ?", id).Update("name", district.Name)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"district": district.Name,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemoveDistrict route struct method
func (gorm *Gorm) DeleteRemoveDistrict(c *gin.Context) {
	var (
		district models.Districts
		result   gin.H
	)

	id := c.Query("id")
	gorm.DB.Delete(&district, id)
	result = gin.H{
		"Message": "Success delete district",
	}
	c.JSON(http.StatusOK, result)
}

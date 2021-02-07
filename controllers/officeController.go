package controllers

import (
	"fmt"
	"net/http"
	"testcasethree-residentAPI/helper"
	"testcasethree-residentAPI/models"

	"github.com/gin-gonic/gin"
)

type officeResponse struct {
	ID          uint   `json:"id"`
	SubDistrict string `json:"sub_district"`
	Office      string `json:"office"`
}

// PostCreateOffice route struct method
func (gorm *Gorm) PostCreateOffice(c *gin.Context) {
	var (
		office models.Offices
		result gin.H
	)

	if err := c.Bind(&office); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Create(&office)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":              office.ID,
				"sub_district_id": office.SubDistrictID,
				"office":          office.Name,
				"created_at":      office.CreatedAt,
				"update_at":       office.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadOffice route func
func (gorm *Gorm) GetReadOffice(c *gin.Context) {
	var (
		office   []models.Offices
		response []officeResponse
		result   gin.H
	)

	gorm.DB.Model(&office).Select("offices.id, offices.name as office, sub_districts.name as subdistrict").Joins("left join sub_districts on sub_districts.id = offices.sub_district_id").Scan(&response)
	if length := len(response); length <= 0 {
		result = helper.ResultAPINilResponse(response, length)
	} else {
		result = helper.ResultAPIResponse(response, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdateOffice route struct method
func (gorm *Gorm) PatchUpdateOffice(c *gin.Context) {
	var (
		office models.Offices
		result gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&office); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Model(&office).Where("id = ?", id).Update("name", office.Name)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"office": office.Name,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemoveOffice route struct method
func (gorm *Gorm) DeleteRemoveOffice(c *gin.Context) {
	var (
		office models.Offices
		result gin.H
	)

	id := c.Query("id")
	gorm.DB.Delete(&office, id)
	result = gin.H{
		"Message": "Success delete office",
	}
	c.JSON(http.StatusOK, result)
}

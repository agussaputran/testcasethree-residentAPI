package controllers

import (
	"fmt"
	"net/http"
	"testcasethree-residentAPI/models"

	"github.com/gin-gonic/gin"
)

type provinceResponse struct {
	ID       uint
	Province string
}

// PostCreateProvince route struct method
func (gorm *Gorm) PostCreateProvince(c *gin.Context) {
	var (
		province models.Provinces
		result   gin.H
	)

	if err := c.Bind(&province); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Create(&province)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":         province.ID,
				"province":   province.Name,
				"created_at": province.CreatedAt,
				"update_at":  province.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadProvince func route
func (gorm *Gorm) GetReadProvince(c *gin.Context) {
	var (
		province []models.Provinces
		// response []provinceResponse
		result gin.H
	)

	gorm.DB.Preload("District").Preload("District.SubDistrict").Preload("District.SubDistrict.Person").Find(&province)
	if length := len(province); length <= 0 {
		result = ResultAPINilResponse(province, length)
	} else {
		result = ResultAPIResponse(province, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdateProvince route struct method
func (gorm *Gorm) PatchUpdateProvince(c *gin.Context) {
	var (
		province models.Provinces
		result   gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&province); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		gorm.DB.Model(&province).Where("id = ?", id).Update("name", province.Name)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"province": province.Name,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemoveProvince route struct method
func (gorm *Gorm) DeleteRemoveProvince(c *gin.Context) {
	var (
		province models.Provinces
		result   gin.H
	)

	id := c.Query("id")
	gorm.DB.Delete(&province, id)
	result = gin.H{
		"Message": "Success delete",
	}
	c.JSON(http.StatusOK, result)
}

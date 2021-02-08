package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"testcasethree-residentAPI/models"

	"github.com/gin-gonic/gin"
)

// UploadSingle func
func (gorm *Gorm) UploadSingle(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	file, _ := c.FormFile("photo")

	if file.Size > 200000 || file.Header["Content-Type"][0] != "image/png" {
		c.JSON(400, gin.H{
			"message": "uploaded file size is too large or not a valid image format jpg/png",
		})
		return
	} else if file.Header["Content-Type"][0] != "image/jpg" {
		c.JSON(400, gin.H{
			"message": "uploaded file size is too large or not a valid image format jpg/png",
		})
		return
	}

	id := c.Query("id")

	reg, err := regexp.Compile("[^a-zA-Z.]+")
	if err != nil {
		log.Fatal(err)
	}

	fileNameLower := strings.ToLower(file.Filename)
	fileNameFmt := strings.ReplaceAll(fileNameLower, " ", "_")
	fileNameFmt2 := reg.ReplaceAllString(fileNameFmt, "")
	fileName := id + "_" + fileNameFmt2
	path := "images/" + fileName

	if err := c.SaveUploadedFile(file, path); err != nil {
		fmt.Println("Terjadi Error", err.Error())
	}

	gorm.DB.Model(&person).Where("id = ?", id).Update("photo_profile_url", path)

	result = gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"localPath": path,
		},
	}
	c.JSON(http.StatusOK, result)
}

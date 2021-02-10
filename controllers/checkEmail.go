package controllers

import (
	"testcasethree-residentAPI/models"

	"gorm.io/gorm"
)

// CheckQueue func
func CheckQueue(db *gorm.DB) {
	var queue models.QueueEmail

	db.Where("handled = ?", false).First(&queue)

}

func SendMail() {

}

func SendMailConfig(to, cc, subject, message string) {

}

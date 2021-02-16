package models

import "gorm.io/gorm"

// QueueEmail model
type QueueEmail struct {
	gorm.Model
	Handled bool
	To      string
	Cc      string
	Subject string
	Message string
}

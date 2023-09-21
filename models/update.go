package models

import (
	"time"
)

type Update struct {
	Id          int       `gorm:"id"`
	JobId       int       `gorm:"job_id"`
	Display     string    `gorm:"display"`
	DateUpdated time.Time `gorm:"date_updated"`
	Notes       string    `gorm:"notes"`
}

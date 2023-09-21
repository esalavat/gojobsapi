package models

import (
	"time"
)

type Job struct {
	Id          int       `gorm:"id"`
	Company     string    `gorm:"company"`
	JobTitle    string    `gorm:"job_title"`
	DateApplied time.Time `gorm:"date_applied"`
	JobUrl      string    `gorm:"job_url"`
	Updates     []Update  `gorm:"foreignKey:job_id"`
}

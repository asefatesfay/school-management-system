package model

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	StudentID uint      `json:"student_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
}

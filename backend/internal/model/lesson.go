package model

import (
	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	SubjectID   uint      `json:"subject_id"`
	ClassID     uint      `json:"class_id"`
}

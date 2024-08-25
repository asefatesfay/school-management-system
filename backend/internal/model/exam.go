package model

import (
	"time"

	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	SubjectID   uint      `json:"subject_id"`
	ClassID     uint      `json:"class_id"`
}

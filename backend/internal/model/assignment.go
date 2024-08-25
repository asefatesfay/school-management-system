package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	SubjectID   uint      `json:"subject_id"`
	ClassID     uint      `json:"class_id"`
}

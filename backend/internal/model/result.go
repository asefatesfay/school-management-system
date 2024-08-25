package model

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	ExamID    uint    `json:"exam_id"`
	Marks     float64 `json:"marks"`
	Grade     string  `json:"grade"`
}

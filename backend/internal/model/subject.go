package model

import (
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	TeacherID   uint   `json:"teacher_id"`
}

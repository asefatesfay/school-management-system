package model

import (
	"gorm.io/gorm"
)

type Parent struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	StudentID   uint   `json:"student_id"`
}

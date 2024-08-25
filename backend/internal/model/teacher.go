package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	HireDate    string `json:"hire_date"`
	Department  string `json:"department"`
}

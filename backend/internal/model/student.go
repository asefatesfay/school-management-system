package model  
  
import (  
    "gorm.io/gorm"  
)  
  
type Student struct {  
    gorm.Model  
    FirstName      string `json:"first_name"`  
    LastName       string `json:"last_name"`  
    BirthDate      string `json:"birth_date"`  
    Email          string `json:"email"`  
    PhoneNumber    string `json:"phone_number"`  
    EnrollmentDate string `json:"enrollment_date"`  
    ClassID        uint   `json:"class_id"`  
}  

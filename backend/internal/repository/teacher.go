package repository

import (
	"school-management/database"
	"school-management/internal/model"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	FindAll() ([]model.Teacher, error)
	FindByID(id uint) (*model.Teacher, error)
	Create(teacher *model.Teacher) error
	Update(teacher *model.Teacher) error
	Delete(id uint) error
}

type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository() TeacherRepository {
	return &teacherRepository{db: database.GetDB()}
}

func (r *teacherRepository) FindAll() ([]model.Teacher, error) {
	var teachers []model.Teacher
	if err := r.db.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *teacherRepository) FindByID(id uint) (*model.Teacher, error) {
	var teacher model.Teacher
	if err := r.db.First(&teacher, id).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *teacherRepository) Create(teacher *model.Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *teacherRepository) Update(teacher *model.Teacher) error {
	return r.db.Save(teacher).Error
}

func (r *teacherRepository) Delete(id uint) error {
	return r.db.Delete(&model.Teacher{}, id).Error
}

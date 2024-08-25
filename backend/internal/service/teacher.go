package service

import (
	"school-management/internal/model"
	"school-management/internal/repository"
)

type TeacherService interface {
	GetAllTeachers() ([]model.Teacher, error)
	GetTeacherByID(id uint) (*model.Teacher, error)
	CreateTeacher(teacher *model.Teacher) error
	UpdateTeacher(teacher *model.Teacher) error
	DeleteTeacher(id uint) error
}

type teacherService struct {
	repository repository.TeacherRepository
}

func NewTeacherService() TeacherService {
	return &teacherService{
		repository: repository.NewTeacherRepository(),
	}
}

func (s *teacherService) GetAllTeachers() ([]model.Teacher, error) {
	return s.repository.FindAll()
}

func (s *teacherService) GetTeacherByID(id uint) (*model.Teacher, error) {
	return s.repository.FindByID(id)
}

func (s *teacherService) CreateTeacher(teacher *model.Teacher) error {
	return s.repository.Create(teacher)
}

func (s *teacherService) UpdateTeacher(teacher *model.Teacher) error {
	return s.repository.Update(teacher)
}

func (s *teacherService) DeleteTeacher(id uint) error {
	return s.repository.Delete(id)
}

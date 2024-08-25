package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"school-management/internal/model"
	"school-management/internal/service"

	"github.com/go-chi/chi/v5"
)

var teacherService = service.NewTeacherService()

// GetAllTeachers retrieves all teachers
func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherService.GetAllTeachers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(teachers)
}

// GetTeacherByID retrieves a teacher by ID
func GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}
	teacher, err := teacherService.GetTeacherByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

// CreateTeacher creates a new teacher
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher model.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := teacherService.CreateTeacher(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)
}

// UpdateTeacher updates an existing teacher
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	var teacher model.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	teacher.ID = uint(id)
	if err := teacherService.UpdateTeacher(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

// DeleteTeacher deletes a teacher by ID
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	if err := teacherService.DeleteTeacher(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

package router

import (
	"net/http"
	"school-management/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Teacher routes
	r.Route("/teachers", func(r chi.Router) {
		r.Get("/", handler.GetAllTeachers)
		r.Get("/{id}", handler.GetTeacherByID)
		r.Post("/", handler.CreateTeacher)
	})

	// Add similar routes for other entities like Students, Parents, Classes, etc.
	// Example for students:
	// r.Route("/students", func(r chi.Router) {
	//     r.Get("/", handler.GetAllStudents)
	//     r.Get("/{id}", handler.GetStudentByID)
	//     r.Post("/", handler.CreateStudent)
	//     r.Put("/{id}", handler.UpdateStudent)
	//     r.Delete("/{id}", handler.DeleteStudent)
	// })

	return r
}

package main

import (
	"log"
	"net/http"
	"school-management/database"
	"school-management/internal/model"
	"school-management/internal/router"
)

func main() {
	// Initialize the database (singleton instance)
	database := database.GetDB()

	// Auto migrate the models (optional)
	database.AutoMigrate(&model.Teacher{}, &model.Student{}, &model.Parent{}, &model.Class{}, &model.Subject{},
		&model.Lesson{}, &model.Exam{}, &model.Assignment{}, &model.Result{}, &model.Attendance{}, &model.Event{}, &model.Message{})

	// Set up the router
	r := router.SetupRouter()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

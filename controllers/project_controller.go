package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"os"
	"your_project/models"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&models.Project{})
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := db.Create(&project); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var updatedProject models.Project
	if err := json.NewDecoder(r.Body).Decode(&updatedProject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Model(&project).Updates(updatedProject)
	json.NewEncoder(w).Encode(project)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	db.Find(&projects)
	json.NewEncoder(w).Encode(projects)
}
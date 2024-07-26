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
	db, err = gorm.Open(os.Getenv("DB_DIALECT"), &gorm.Config{
        DSN: os.Getenv("DB_CONNECTION_STRING"),
    })
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	if err := db.AutoMigrate(&models.Project{}); err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if result := db.Create(&project); result.Error != nil {
		http.Error(w, "Error creating project: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(project); err != nil {
        http.Error(w, "Error encoding project: "+err.Error(), http.StatusInternalServerError)
    }
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		http.Error(w, "Project not found: "+err.Error(), http.StatusNotFound)
		return
	}

	var updatedProject models.Project
	if err := json.NewDecoder(r.Body).Decode(&updatedProject); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Model(&project).Updates(updatedProject).Error; err != nil {
	    http.Error(w, "Error updating project: "+err.Error(), http.StatusInternalServerError)
	    return
	}
	
	if err := json.NewEncoder(w).Encode(project); err != nil {
        http.Error(w, "Error encoding project: "+err.Error(), http.StatusInternalServerError)
    }
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		http.Error(w, "Project not found: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		http.Error(w, "Error deleting project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		http.Error(w, "Project not found: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(project); err != nil {
        http.Error(w, "Error encoding project: "+err.Error(), http.StatusInternalServerError)
    }
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	if result := db.Find(&projects); result.Error != nil {
	    http.Error(w, "Error fetching projects: "+result.Error.Error(), http.StatusInternalServerError)
	    return
	}

	if err := json.NewEncoder(w).Encode(projects); err != nil {
        http.Error(w, "Error encoding projects: "+err.Error(), http.StatusInternalServerError)
    }
}
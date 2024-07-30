package controllers

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "net/http"
    "os"
    "sync"
    "your_project/models"
)

var db *gorm.DB
var projectCache sync.Map 
var projectsCache []models.Project 
var projectsCacheValid bool 

func init() {
    var err error
	db, err = gorm.Open(gorm.Open(os.Getenv("DB_DIALECT"), &gorm.Config{}), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	if err := db.AutoMigrate(&models.Project{}); err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
    projectsCacheValid = false
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
    projectsCacheValid = false
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
    projectsCacheValid = false
}

func GetProject(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if cachedProject, ok := projectCache.Load(id); ok {
        if err := json.NewEncoder(w).Encode(cachedProject); err != nil {
            http.Error(w, "Error encoding project: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }

    var project models.Project
    if err := db.First(&project, id).Error; err != nil {
        http.Error(w, "Project not found: "+err.Error(), http.StatusNotFound)
        return
    }

    projectCache.Store(id, project)

    if err := json.NewEncoder(w).Encode(project); err != nil {
        http.Error(w, "Error encoding project: "+err.Error(), http.StatusInternalServerError)
    }
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
    if projectsCacheValid {
        if err := json.NewEncoder(w).Encode(projectsCache); err != nil {
            http.Error(w, "Error encoding projects: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }

    var projects []models.Project
    if result := db.Find(&projects); result.Error != nil {
        http.Error(w, "Error fetching projects: "+result.Error.Error(), http.StatusInternalServerError)
        return
    }

    projectsCache = projects
    projectsCacheValid = true

    if err := json.NewEncoder(w).Encode(projects); err != nil {
        http.Error(w, "Error encoding projects: "+err.Error(), http.StatusInternalServerError)
    }
}

func ClearProjectsCache() {
    projectCache = sync.Map{} 
    projectsCacheValid = false
}
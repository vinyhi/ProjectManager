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
var projectCache sync.Map // Initialize a new sync.Map for caching projects
var projectsCache []models.Project // Slice to cache all projects, simplification for demonstration
var projectsCacheValid bool // Indicates if the projectsCache is currently valid

func init() {
    var err error
	db, err = gorm.Open(os.Getenv("DB_DIALECT"), &gorm.Config{
		DSN: os.Getenv("DB_CONNECTION_STRING"),
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	if err := db.AutoMigrate(&models.Project{}); err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
    // Same implementation, but remember to invalidate the cache
    // Add a line to invalidate the projectsCache upon successful project creation
    projectsCacheValid = false
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
    // Same implementation, but remember to invalidate the cache for the updated project
    // Invalidate both specific project cache and all projects cache
    projectsCacheValid = false
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
    // Same implementation, but invalidate the cache for the deleted project
    projectsCacheValid = false
}

func GetProject(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    // Check if the project is in cache first
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

    // Cache the project before sending the response
    projectCache.Store(id, project)

    if err := json.NewEncoder(w).Encode(project); err != nil {
        http.Error(w, "Error encoding project: "+err.Error(), http.StatusInternalServerError)
    }
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
    // Check if we already have a valid cache of all projects
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

    // Update the projectsCache and mark it as valid
    projectsCache = projects
    projectsCacheValid = true

    if err := json.NewEncoder(w).Encode(projects); err != nil {
        http.Error(w, "Error encoding projects: "+err.Error(), http.StatusInternalServerError)
    }
}
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
	"github.com/julienschmidt/httprouter"
)

// ListProjects returns a list of all the projects created by members
func ListProjects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var projects []models.Project
	db.Find(&projects)

	respBody, err := json.MarshalIndent(projects, "", " ")
	if err != nil {
		log.Fatalf("Error listing projects: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

// ShowProject returns the data for a project with given id
func ShowProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var project models.Project
	db.First(&project, ps.ByName("id"))

	// If the project was found, its ID must be different than 0
	if project.ID != 0 {
		respBody, err := json.MarshalIndent(project, "", " ")
		if err != nil {
			log.Fatalf("Error showing project: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusOK)
		return
	}

	utils.JSONMessage(w, "Project not found", http.StatusNotFound)
}

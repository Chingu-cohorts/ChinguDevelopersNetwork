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

// CreateProject saves a new project for a given user
func CreateProject(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// We need to get the user id from the token
	// then we can proceed to load the user variable
	// with the data for the user from the database
	// and then we can append the new project

	var project models.Project

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)
	if err != nil {
		utils.JSONMessage(w, "Wrong project body", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Error reading token while creating post", http.StatusInternalServerError)
		return
	}

	var user models.User
	// Find the user in the db
	db.First(&user, tokenData.User.ID)

	// If we found the user we can append the project
	if user.ID != 0 {
		db.Model(&user).Association("Projects").Append(&project)
		utils.JSONMessage(w, "Project saved successfully", http.StatusCreated)
		return
	}

	utils.JSONMessage(w, "User not found", http.StatusNotFound)
}

// AppendUser adds a new user to the project
func AppendUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// ToDo: I have to create an struct which includes both
	// the project id and the user id (or some sort of similar)
	// unique identifiers

	// We have to find the project in the database so we can append an user to it
	var project models.Project
	db.First(&project, ps.ByName("projectID"))

	// If we found the project then we can proceed
	if project.ID != 0 {
		// Do stuff
	}

	utils.JSONMessage(w, "Could not append project to user", http.StatusInternalServerError)
}

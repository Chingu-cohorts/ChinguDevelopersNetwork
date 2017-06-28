package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chingu-cohorts/ChinguCentral/models"
	"github.com/Chingu-cohorts/ChinguCentral/utils"
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

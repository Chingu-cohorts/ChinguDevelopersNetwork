package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Chingu-cohorts/ChinguCentral/models"
	"github.com/Chingu-cohorts/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
)

// ListCohorts returns a list of all the cohorts in the database
func ListCohorts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var cohorts []models.Cohort
	db.Preload("Users.Cohort").Find(&cohorts)

	respBody, err := json.MarshalIndent(cohorts, "", " ")
	if err != nil {
		log.Fatalf("Error listing cohorts: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

// ShowCohort return the data for a cohort with given id
func ShowCohort(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var cohort models.Cohort
	// We have to modify the string as the query must match exactly the name of the cohort
	lowerCaseCohortName := strings.ToLower(ps.ByName("name"))
	validCohortName := strings.Title(lowerCaseCohortName)
	db.Where("name = ?", validCohortName).Preload("Users.Cohort").First(&cohort)

	if cohort.ID != 0 {
		respBody, err := json.MarshalIndent(cohort, "", " ")
		if err != nil {
			log.Fatalf("Error showing cohort: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusOK)
		return
	}

	utils.JSONMessage(w, "Cohort not found", http.StatusNotFound)
}

// CreateCohort saves a new cohort to the database
func CreateCohort(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var cohort models.Cohort

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cohort)
	if err != nil {
		utils.JSONMessage(w, "Wrong body", http.StatusBadRequest)
		return
	}

	// Verify name and description as both are required
	if cohort.Name != "" && cohort.Description != "" {
		// Verify if there's a cohort with that name already
		var existingCohort models.Cohort
		db.Where("name = ?", cohort.Name).First(&existingCohort)

		// Cohort with that name already exists, don't create it
		if existingCohort.ID != 0 {
			utils.JSONMessage(w, "Cohort with given name already exists", http.StatusBadRequest)
			return
		}

		// Cohort with given name doesn't exist, go ahead and create it
		db.Create(&cohort)

	} else {
		utils.JSONMessage(w, "Name and Description are required", http.StatusBadRequest)
		return
	}

	respBody, err := json.MarshalIndent(cohort, "", " ")
	if err != nil {
		log.Fatalf("Error returning created cohort: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusCreated)
}

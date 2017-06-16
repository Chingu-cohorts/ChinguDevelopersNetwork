package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chingu-cohorts/ChinguCentral/models"
	"github.com/Chingu-cohorts/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
)

// ListCohorts returns a list of all the cohorts in the database
func ListCohorts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var cohorts []models.Cohort
	db.Preload("Users").Find(&cohorts)

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
	db.Preload("Users").First(&cohort, ps.ByName("id"))

	if cohort.ID != 0 {
		respBody, err := json.MarshalIndent(cohort, "", " ")
		if err != nil {
			log.Fatalf("Error showing cohort: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusOK)
	}

	utils.JSONMessage(w, "Cohort not found", http.StatusNotFound)
}

// CreateCohort saves a new cohort to the database
// ToDo: Fix message when a cohort already exists
func CreateCohort(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var cohort models.Cohort

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cohort)
	if err != nil {
		utils.JSONMessage(w, "Wrong body", http.StatusBadRequest)
	}

	if cohort.Name != "" && cohort.Description != "" {
		db.Create(&cohort)
	} else {
		utils.JSONMessage(w, "Name and Description are required", http.StatusBadRequest)
	}

	respBody, err := json.MarshalIndent(cohort, "", " ")
	if err != nil {
		log.Fatalf("Error returning created cohort: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusCreated)
}

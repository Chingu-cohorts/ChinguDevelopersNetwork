package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Oxyrus/ChinguCentral/models"
	"github.com/Oxyrus/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
)

// AllCohorts returns a list of the cohorts in the database
func AllCohorts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var cohorts []models.Cohort
	db.Find(&cohorts)
	respBody, err := json.MarshalIndent(cohorts, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	utils.ResponseWithJSON(w, respBody, http.StatusOK)
}

// CohortByID finds a cohort given a parameter with its id
func CohortByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	id := ps.ByName("id")

	var cohort models.Cohort

	db.First(&cohort, id)

	if cohort.ID != 0 {
		respBody, err := json.MarshalIndent(cohort, "", " ")
		if err != nil {
			log.Fatal(err)
			return
		}
		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	} else {
		utils.ErrorWithJSON(w, "Cohort not found", http.StatusNotFound)
	}
}

// CreateCohort saves a cohort to the database
func CreateCohort(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var cohort models.Cohort

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cohort)
	if err != nil {
		utils.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	if cohort.Name != "" && cohort.Description != "" {
		db.Create(&cohort)
	} else {
		utils.ErrorWithJSON(w, "Name and Description are required", http.StatusBadRequest)
		return
	}

	respBody, err := json.MarshalIndent(cohort, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	utils.ResponseWithJSON(w, respBody, http.StatusCreated)
}

// DeleteCohort finds a cohort by its name and removes it from the database
func DeleteCohort(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	id := ps.ByName("id")

	var cohort models.Cohort
	db.First(&cohort, id)

	if cohort.ID != 0 {
		db.Delete(&cohort)
		w.WriteHeader(http.StatusNoContent)
	} else {
		utils.ErrorWithJSON(w, "Cohort not found", http.StatusNotFound)
	}
}

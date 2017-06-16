package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chingu-cohorts/ChinguCentral/models"
	"github.com/Chingu-cohorts/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
)

// ListUsers returns a list of all the users in the database
func ListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var users []models.User
	db.Preload("Cohort").Find(&users)

	respBody, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatalf("Error listing the users: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

// ShowUser returns the data for an user with given id
func ShowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var user models.User
	db.Preload("Cohort").First(&user, ps.ByName("id"))

	if user.ID != 0 {
		respBody, err := json.MarshalIndent(user, "", " ")
		if err != nil {
			log.Fatalf("Error showing user: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusOK)
	}

	utils.JSONMessage(w, "User not found", http.StatusNotFound)
}

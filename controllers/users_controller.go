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

// CreateUser registers a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.JSONMessage(w, "Wrong body", http.StatusBadRequest)
		return
	}

	if user.Username != "" && user.Email != "" {
		// Used to check for an user with the same email
		var existingUserWithEmail models.User
		db.Where("email = ?", user.Email).First(&existingUserWithEmail)

		// Used to check for an user with the same username
		var existingUserWithUsername models.User
		db.Where("username = ?", user.Username).First(&existingUserWithUsername)

		// There's an user with the same email and username
		if existingUserWithEmail.ID != 0 && existingUserWithUsername.ID != 0 {
			utils.JSONMessage(w, "Username and email already taken", http.StatusOK)
			return
		}

		// Email already registered
		if existingUserWithEmail.ID != 0 {
			utils.JSONMessage(w, "Email already registered", http.StatusOK)
			return
		}

		// Username already registered
		if existingUserWithUsername.ID != 0 {
			utils.JSONMessage(w, "Username already registered", http.StatusOK)
			return
		}

		// User was successfully created, format to JSON and response
		db.Create(&user)

		respBody, err := json.MarshalIndent(user, "", " ")
		if err != nil {
			log.Fatalf("Error formatting created user: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusCreated)
		return
	}

	utils.JSONMessage(w, "Email and Username are required", http.StatusBadRequest)
}

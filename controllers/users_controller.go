package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

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
	// Username must match exactly the one stored
	lowercaseParams := strings.ToLower(ps.ByName("username"))
	db.Where("username = ?", lowercaseParams).Preload("Cohort").First(&user)
	db.Model(&user).Association("Projects").Find(&user.Projects)

	if user.ID != 0 {
		respBody, err := json.MarshalIndent(user, "", " ")
		if err != nil {
			log.Fatalf("Error showing user: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusOK)
		return
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

	if user.Username != "" && user.Email != "" && user.Password != "" {
		// Only lowercase usernames allowed
		user.Username = strings.ToLower(user.Username)

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

		// We can create the user
		// Start by hashing the password
		password := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Error hashing password: %v", err)
		}

		// Assign the user the hashed password
		user.EncryptedPassword = string(hashedPassword)

		// Save the user to the database
		db.Create(&user)

		respBody, err := json.MarshalIndent(user, "", " ")
		if err != nil {
			log.Fatalf("Error formatting created user: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusCreated)
		return
	}

	utils.JSONMessage(w, "Email, username and password are required", http.StatusBadRequest)
}

// Login verifies the credentials and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// The user that comes along with the POST request
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.JSONMessage(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// savedUser is the user found in the database
	var savedUser models.User
	username := user.Username

	db.Where("username = ?", username).First(&savedUser)

	if savedUser.ID == 0 {
		utils.JSONMessage(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(savedUser.EncryptedPassword), []byte(user.Password))
	if err != nil {
		utils.JSONMessage(w, "Wrong password", http.StatusBadRequest)
		return
	}

	utils.JSONMessage(w, "Correct password", http.StatusOK)
}

// DeleteUser finds and user by its username and deletes it
func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var user models.User
	db.Where("username = ?", ps.ByName("username")).First(&user)

	if user.ID != 0 {
		db.Delete(&user)
		utils.JSONMessage(w, "User deleted", http.StatusOK)
		return
	}

	utils.JSONMessage(w, "User not found", http.StatusOK)
}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
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
	// ORM things, don't ask
	db.Where("username = ?", lowercaseParams).Preload("Cohort").First(&user)
	db.Model(&user).Association("Projects").Find(&user.Projects)

	// If the user exists, its ID must be different than 0
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

	// If there's valid data, we can start working
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

	// If we didn't find the user
	if savedUser.ID == 0 {
		utils.JSONMessage(w, "User not found", http.StatusNotFound)
		return
	}

	// Compares the saved hash with the password sent in the request
	err = bcrypt.CompareHashAndPassword([]byte(savedUser.EncryptedPassword), []byte(user.Password))
	if err != nil {
		utils.JSONMessage(w, "Wrong password", http.StatusBadRequest)
		return
	}

	// Generate JWT and send it
	token := utils.GenerateJWT(savedUser)
	result := models.ResponseToken{Token: token}

	respBody, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatalf("Error marshalizing token: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

// DeleteUser finds and user by its username and deletes it
func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var user models.User
	db.Where("username = ?", ps.ByName("username")).First(&user)

	// If we found the user then we can proceed
	if user.ID != 0 {
		db.Delete(&user)
		utils.JSONMessage(w, "User deleted", http.StatusOK)
		return
	}

	utils.JSONMessage(w, "User not found", http.StatusOK)
}

// UpdateUser saves the new data for a given user
func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// The user that comes along the PUT request
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.JSONMessage(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Get the user id from the JWT
	token := r.Header.Get("Authorization")

	// Read the data from the JWT
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Error ocurred reading the token", http.StatusInternalServerError)
		return
	}

	user.ID = tokenData.User.ID

	// We want to make sure a non-admin user
	// is not making himself an admin
	var savedUser models.User

	// We lookup for the user and write the data to
	// the savedUser variable
	db.First(&savedUser, user.ID)

	// If the user is not an admin already, force it
	// to have a false value, which means admin can
	// only be given through database access
	if !savedUser.IsAdmin {
		user.IsAdmin = false
	}

	db.Model(&user).Updates(user)

	utils.JSONMessage(w, "User updated successfully", http.StatusAccepted)
}

// CurrentUser returns the data for the logged in user
func CurrentUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// First we get the JWT token from the request
	token := r.Header.Get("Authorization")

	// Read the data from the JWT
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Something ocurred while reading the token", http.StatusInternalServerError)
		return
	}

	// The data contained in the token is old, all we need
	// is the id, so we can look up for the user in the database
	// and return the data
	userID := tokenData.User.ID

	var user models.User

	db.Where("id = ?", userID).Preload("Cohort").First(&user)
	db.Model(&user).Association("Projects").Find(&user.Projects)

	respBody, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatalf("Error in current user: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

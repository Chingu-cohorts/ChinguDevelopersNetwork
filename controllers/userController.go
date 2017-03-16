package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Oxyrus/ChinguCentral/models"
	"github.com/Oxyrus/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
)

// AllUsers returns a list of all the records in the database
func AllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var users []models.User
	db.Find(&users)
	respBody, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	utils.ResponseWithJSON(w, respBody, http.StatusOK)
}

// UserByID returns the data in the database for the request parameter
func UserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var user models.User
	id := ps.ByName("id")

	db.First(&user, id)
	if user.ID != 0 {
		respBody, err := json.MarshalIndent(user, "", " ")
		if err != nil {
			log.Fatal(err)
			return
		}
		utils.ResponseWithJSON(w, respBody, http.StatusOK)
		return
	}

	utils.ErrorWithJSON(w, "User not found", http.StatusNotFound)
}

// CreateUser registers an user in the database
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.ErrorWithJSON(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if user.Email != "" && user.Name != "" && user.Username != "" {
		db.Create(&user)
	} else {
		utils.ErrorWithJSON(w, "Email, Name and Username are required", http.StatusBadRequest)
		return
	}

	respBody, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	utils.ResponseWithJSON(w, respBody, http.StatusCreated)
}

// DeleteUser finds an user by its name and removes it from the database
func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	id := ps.ByName("id")

	var user models.User
	db.Find(&user, id)
	if user.ID != 0 {
		db.Delete(&user)
		w.WriteHeader(http.StatusNoContent)
	} else {
		utils.ErrorWithJSON(w, "User not found", http.StatusNotFound)
	}
}

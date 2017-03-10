package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Oxyrus/ChinguCentral/models"
	"github.com/Oxyrus/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AllUsers returns a list of all the records in the database
func AllUsers(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("caronte").C("users")

		var users []models.User

		err := c.Find(nil).All(&users)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		}

		respBody, err := json.MarshalIndent(users, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// UserByName returns the data in the database for the request parameter
func UserByName(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		name := ps.ByName("name")
		var user models.User

		c := session.DB("caronte").C("users")

		err := c.Find(bson.M{"name": name}).One(&user)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("There was a problem with the request", err)
			return
		}

		if user.Name == "" {
			utils.ErrorWithJSON(w, "User not found", http.StatusNotFound)
			return
		}

		respBody, err := json.MarshalIndent(user, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// CreateUser registers an user in the database
func CreateUser(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		var user models.User

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			utils.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		user.CreatedAt = time.Now()
		c := session.DB("caronte").C("users")

		err = c.Insert(user)
		if err != nil {
			if mgo.IsDup(err) {
				utils.ErrorWithJSON(w, "User already exists", http.StatusBadRequest)
				return
			}
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed to insert user: ", err)
			return
		}

		respBody, err := json.MarshalIndent(user, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusCreated)
	}
}

// UpdateUser finds an user by its name and updates its information
func UpdateUser(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		var user models.User
		name := ps.ByName("name")

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			utils.ErrorWithJSON(w, "Wrong body", http.StatusBadRequest)
			return
		}

		c := session.DB("caronte").C("users")

		err = c.Update(bson.M{"name": name}, user)
		if err != nil {
			switch err {
			default:
				utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed to update user")
				return
			case mgo.ErrNotFound:
				utils.ErrorWithJSON(w, "User not found", http.StatusNotFound)
				log.Println("User not found: ", err)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteUser finds an user by its name and removes it from the database
func DeleteUser(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		name := ps.ByName("name")
		c := session.DB("caronte").C("users")

		err := c.Remove(bson.M{"name": name})
		if err != nil {
			switch err {
			default:
				utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				return
			case mgo.ErrNotFound:
				utils.ErrorWithJSON(w, "User not found", http.StatusNotFound)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

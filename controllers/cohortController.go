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

// AllCohorts returns a list of the cohorts in the database
func AllCohorts(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("caronte").C("cohorts")

		var cohorts []models.Cohort
		err := c.Find(nil).All(&cohorts)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		}

		respBody, err := json.MarshalIndent(cohorts, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// CohortByName finds a cohort given a parameter with its name
func CohortByName(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		name := ps.ByName("name")
		var cohort models.Cohort

		c := session.DB("caronte").C("cohorts")

		err := c.Find(bson.M{"name": name}).One(&cohort)
		if err != nil {
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("There was an error with the request", err)
			return
		}

		if cohort.Name == "" {
			utils.ErrorWithJSON(w, "Cohort not found", http.StatusNotFound)
			return
		}

		respBody, err := json.MarshalIndent(cohort, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// CreateCohort saves a cohort to the database
func CreateCohort(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		var cohort models.Cohort

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&cohort)
		if err != nil {
			utils.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		cohort.CreatedAt = time.Now()

		c := session.DB("caronte").C("cohorts")

		err = c.Insert(cohort)
		if err != nil {
			if mgo.IsDup(err) {
				utils.ErrorWithJSON(w, "Cohort already exists", http.StatusBadRequest)
				return
			}
			utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			return
		}

		respBody, err := json.MarshalIndent(cohort, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		utils.ResponseWithJSON(w, respBody, http.StatusCreated)
	}
}

// UpdateCohort finds a cohort by its name and updates it
func UpdateCohort(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		var cohort models.Cohort

		name := ps.ByName("name")

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&cohort)
		if err != nil {
			utils.ErrorWithJSON(w, "Invalid body", http.StatusBadRequest)
			return
		}

		c := session.DB("caronte").C("cohorts")

		err = c.Update(bson.M{"name": name}, &cohort)
		if err != nil {
			switch err {
			default:
				utils.ErrorWithJSON(w, "Database Error", http.StatusInternalServerError)
				log.Println("Failed to update cohort: ", err)
				return
			case mgo.ErrNotFound:
				utils.ErrorWithJSON(w, "Cohort not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteCohort finds a cohort by its name and removes it from the database
func DeleteCohort(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		name := ps.ByName("name")
		c := session.DB("caronte").C("cohorts")

		err := c.Remove(bson.M{"name": name})
		if err != nil {
			switch err {
			default:
				utils.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				return
			case mgo.ErrNotFound:
				utils.ErrorWithJSON(w, "Book not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

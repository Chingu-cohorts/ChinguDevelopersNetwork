package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Oxyrus/ChinguCentral/models"
	"github.com/Oxyrus/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

// AllUsers returns a list of all the records in the database
func AllUsers(s *mgo.Session) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("chingucentral").C("users")

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

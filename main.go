package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Oxyrus/ChinguCentral/controllers"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	router := httprouter.New()

	// router.GET("/", rootHandler(session))
	router.GET("/cohorts", controllers.AllCohorts(session))
	router.GET("/cohorts/:name", controllers.CohortByName(session))
	router.POST("/cohorts", controllers.CreateCohort(session))
	router.GET("/users", controllers.AllUsers(session))
	router.POST("/users", controllers.CreateUser(session))
	router.GET("/users/:name", controllers.UserByName(session))

	n := negroni.Classic()
	n.UseHandler(router)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

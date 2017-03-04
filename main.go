package main

import (
	"log"
	"net/http"

	"github.com/Oxyrus/ChinguCentral/controllers"
	"github.com/julienschmidt/httprouter"
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
	// router.GET("/cohorts", cohortsHandler(session))
	// router.GET("/cohorts/:name", cohortHandler(session))
	router.GET("/users", controllers.AllUsers(session))
	router.POST("/users", controllers.CreateUser(session))
	router.GET("/users/:name", controllers.UserByName(session))
	// router.GET("/users/:name", userHandler(session))

	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Oxyrus/ChinguCentral/controllers"
	"github.com/Oxyrus/ChinguCentral/models"
	"github.com/Oxyrus/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func init() {
	db := utils.InitDB()
	db.AutoMigrate(&models.Cohort{}, &models.User{})
}

func main() {
	router := httprouter.New()

	// Cohort routes
	router.GET("/cohorts", controllers.AllCohorts)
	router.GET("/cohorts/:id", controllers.CohortByID)
	router.POST("/cohorts", controllers.CreateCohort)
	// router.PUT("/cohorts/:name", controllers.UpdateCohort)
	router.DELETE("/cohorts/:id", controllers.DeleteCohort)

	// User routes
	router.GET("/users", controllers.AllUsers)
	router.GET("/users/:id", controllers.UserByID)
	router.POST("/users", controllers.CreateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

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

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Chingu-cohorts/ChinguCentral/controllers"
	"github.com/Chingu-cohorts/ChinguCentral/models"
	"github.com/Chingu-cohorts/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func init() {
	// When initializing the application, we must run the migrations
	db := utils.InitDB()
	defer db.Close()

	db.AutoMigrate(&models.Cohort{}, &models.User{}, &models.Project{})

	cohorts, err := utils.LoadCohortSeed("cohorts.json")
	if err != nil {
		log.Fatalf("Something went wrong reading the cohorts file: %v", err)
	}

	// Iterate over cohorts to save them
	for _, cohort := range cohorts.Cohorts {
		db.Create(&cohort)
	}
}

func main() {
	// Load the config file
	config, err := utils.LoadSettings("config.json")
	if err != nil {
		log.Fatalf("Something went wrong reading the config file: %v", err)
	}

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
	})

	// Instantiate router
	r := httprouter.New()

	r.GET("/api/cohorts", controllers.ListCohorts)
	r.GET("/api/cohorts/:name", controllers.ShowCohort)
	r.POST("/api/cohorts", controllers.CreateCohort)

	r.GET("/api/users", controllers.ListUsers)
	r.GET("/api/users/:username", controllers.ShowUser)
	r.POST("/api/users", controllers.CreateUser)
	r.POST("/api/users/login", controllers.Login)
	r.DELETE("/api/users/:username", controllers.DeleteUser)

	r.GET("/api/projects", controllers.ListProjects)

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(r)

	// Configure server
	s := &http.Server{
		Addr:           config.Port,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

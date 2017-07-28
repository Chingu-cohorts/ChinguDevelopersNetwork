package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/controllers"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	negroniprometheus "github.com/zbindenren/negroni-prometheus"
)

func init() {
	// When initializing the application, we must run the migrations
	db := utils.InitDB()
	defer db.Close()

	db.AutoMigrate(&models.Cohort{}, &models.User{}, &models.Project{}, &models.Aptitude{}, &models.Post{}, &models.Comment{})

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
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Authorization", "Content-Type"},
		Debug:            config.Debug,
	})

	m := negroniprometheus.NewMiddleware("ChinguDevelopersNetwork")

	// Instantiate router
	r := httprouter.New()

	r.Handler("GET", "/metrics", prometheus.Handler())

	r.GET("/api/cohorts", controllers.ListCohorts)
	r.GET("/api/cohorts/:name", controllers.ShowCohort)
	r.POST("/api/cohorts", utils.AuthRequest(controllers.CreateCohort))

	r.GET("/api/users", controllers.ListUsers)
	r.GET("/api/users/:username", controllers.ShowUser)
	r.POST("/api/users", controllers.CreateUser)
	r.PUT("/api/users", utils.AuthRequest(controllers.UpdateUser))
	r.POST("/api/users/login", controllers.Login)
	r.DELETE("/api/users/:username", utils.AuthRequest(controllers.DeleteUser))
	r.GET("/api/currentuser", utils.AuthRequest(controllers.CurrentUser))

	r.GET("/api/projects", controllers.ListProjects)
	r.GET("/api/projects/:id", controllers.ShowProject)
	r.POST("/api/projects", utils.AuthRequest(controllers.CreateProject))

	r.GET("/api/posts", controllers.ListPosts)
	r.GET("/api/posts/:postID", controllers.ShowPost)
	r.POST("/api/posts", utils.AuthRequest(controllers.CreatePost))
	r.DELETE("/api/posts/:postID", utils.AuthRequest(controllers.DeletePost))

	r.POST("/api/posts/:postID/comments", utils.AuthRequest(controllers.CreateComment))

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	// Cors
	n.Use(c)
	// Prometheus
	n.Use(m)
	n.UseHandler(r)

	// Temporary port to test on heroku
	// If there is no env variable called port
	// then use whatever is inside the config file
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = config.Port
	}

	// Configure server
	s := &http.Server{
		Addr:           port,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

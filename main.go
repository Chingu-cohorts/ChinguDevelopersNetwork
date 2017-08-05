package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/controllers"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	negroniprometheus "github.com/zbindenren/negroni-prometheus"
)

func main() {
	// Load the config file
	config, err := utils.LoadSettings("config.json")
	if err != nil {
		logrus.WithError(err)
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

	// For some nice-to-have metrics
	m := negroniprometheus.NewMiddleware("ChinguDevelopersNetwork")

	// Instantiate router
	r := httprouter.New()

	// Prometheus route
	r.Handler("GET", "/metrics", prometheus.Handler())

	// Cohorts routes
	r.GET("/api/cohorts", controllers.ListCohorts)
	r.GET("/api/cohorts/:name", controllers.ShowCohort)
	r.POST("/api/cohorts", utils.AuthRequest(controllers.CreateCohort))

	// User routes
	r.GET("/api/users", controllers.ListUsers)
	r.GET("/api/users/:username", controllers.ShowUser)
	r.POST("/api/users", controllers.CreateUser)
	r.PUT("/api/users", utils.AuthRequest(controllers.UpdateUser))
	r.POST("/api/users/login", controllers.Login)
	r.DELETE("/api/users/:username", utils.AuthRequest(controllers.DeleteUser))
	r.GET("/api/currentuser", utils.AuthRequest(controllers.CurrentUser))

	// Projects routes
	r.GET("/api/projects", controllers.ListProjects)
	r.GET("/api/projects/:id", controllers.ShowProject)
	r.POST("/api/projects", utils.AuthRequest(controllers.CreateProject))

	// Posts routes
	r.GET("/api/posts", controllers.ListPosts)
	r.GET("/api/posts/:postID", controllers.ShowPost)
	r.POST("/api/posts", utils.AuthRequest(controllers.CreatePost))
	r.PUT("/api/posts/:postID", utils.AuthRequest(controllers.UpdatePost))
	r.DELETE("/api/posts/:postID", utils.AuthRequest(controllers.DeletePost))

	// Comments routes
	r.POST("/api/posts/:postID/comments", utils.AuthRequest(controllers.CreateComment))
	r.DELETE("/api/posts/:postID/comments/:commentID", utils.AuthRequest(controllers.DeleteComment))

	// Panic recover, request/response logger, static file serving
	n := negroni.Classic()
	// Gzip responses
	n.Use(gzip.Gzip(gzip.DefaultCompression))

	// Use Cors
	n.Use(c)

	// Use Prometheus
	n.Use(m)

	// Use router instance
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

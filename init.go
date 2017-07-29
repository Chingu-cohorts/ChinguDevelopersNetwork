package main

import (
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func init() {
	// To use cool colors in windows
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	// When initializing the application, we must run the migrations
	db := utils.InitDB()
	defer db.Close()

	db.AutoMigrate(&models.Cohort{}, &models.User{}, &models.Project{}, &models.Aptitude{}, &models.Post{}, &models.Comment{})

	// Load cohorts from file
	cohorts, err := utils.LoadCohortSeed("cohorts.json")
	if err != nil {
		logrus.Panic("Could not read the cohorts file")
	}

	// Iterate over cohorts to save them
	for _, cohort := range cohorts.Cohorts {
		db.Create(&cohort)

		logrus.WithFields(logrus.Fields{
			"cohort": &cohort,
		}).Info("Saving cohort")
	}

	// Load aptitudes from file
	aptitudes, err := utils.LoadAptitudeSeed("aptitudes.json")
	if err != nil {
		logrus.Panic("Could not read the aptitudes file")
	}

	// Iterate over aptitudes to save them
	for _, aptitude := range aptitudes.Aptitudes {
		db.Create(&aptitude)

		logrus.WithFields(logrus.Fields{
			"aptitude": &aptitude,
		}).Info("Saving aptitude")
	}
}

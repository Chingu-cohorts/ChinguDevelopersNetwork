package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
)

// CohortSeed represents the data contained in the cohorts.json file
type CohortSeed struct {
	Cohorts []models.Cohort `json:"cohorts"`
}

// LoadCohortSeed loads the seed for the database
func LoadCohortSeed(filename string) (CohortSeed, error) {
	var cohorts CohortSeed
	cohortsFile, err := os.Open(filename)
	defer cohortsFile.Close()
	if err != nil {
		log.Fatalf("Error reading the cohorts file: %v", err)
		return cohorts, err
	}

	jsonParser := json.NewDecoder(cohortsFile)
	err = jsonParser.Decode(&cohorts)

	return cohorts, err
}

// AptitudeSeed represents the data contained in the aptitudes.json file
type AptitudeSeed struct {
	Aptitudes []models.Aptitude `json:"aptitudes"`
}

// LoadAptitudeSeed loads the data from the aptitudes.json file and
// returns a new struct
func LoadAptitudeSeed(filename string) (AptitudeSeed, error) {
	var aptitudes AptitudeSeed
	aptitudesFile, err := os.Open(filename)
	defer aptitudesFile.Close()
	if err != nil {
		log.Fatalf("Error reading the aptitudes file: %v", err)
		return aptitudes, err
	}

	jsonParser := json.NewDecoder(aptitudesFile)
	err = jsonParser.Decode(&aptitudes)

	return aptitudes, err
}

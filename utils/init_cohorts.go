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

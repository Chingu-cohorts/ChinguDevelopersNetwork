package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/controllers"
	"github.com/julienschmidt/httprouter"
)

// I still have to look for how to fix
// the file path so the tests pass

func TestListCohorts(t *testing.T) {
	handler := controllers.ListCohorts
	router := httprouter.New()

	router.GET("/api/cohorts", handler)
	req, _ := http.NewRequest("GET", "/api/cohorts", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}

package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// I still have to look for how to fix
// the file path so the tests pass

func TestListCohorts(t *testing.T) {
	handler := ListCohorts
	router := httprouter.New()

	router.GET("/api/cohorts", handler)
	req, _ := http.NewRequest("GET", "/api/cohorts", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}

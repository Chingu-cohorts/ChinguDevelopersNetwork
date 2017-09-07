package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
	"github.com/julienschmidt/httprouter"
)

// ListPostsReports returns a list of all non closed reports for posts
func ListPostsReports(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var postReports []models.PostReport
	db.Where("is_closed = ?", false).Find(&postReports)

	respBody, err := json.MarshalIndent(postReports, "", " ")
	if err != nil {
		log.Fatalf("Error listing posts reports: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

// ListCommentsReports returns a list for all non closed reports
func ListCommentsReports(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var commentReports []models.CommentReport
	db.Where("is_closed = ?", false).Find(&commentReports)

	respBody, err := json.MarshalIndent(commentReports, "", " ")
	if err != nil {
		log.Fatalf("Error listing comments reports: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

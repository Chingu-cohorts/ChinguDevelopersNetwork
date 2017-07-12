package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Chingu-cohorts/ChinguCentral/models"
	"github.com/Chingu-cohorts/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
)

// CreateComment saves a new comment to the database for a given post
func CreateComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var comment models.Comment

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		utils.JSONMessage(w, "Wrong body", http.StatusBadRequest)
		return
	}

	// Get the user id from the token
	token := r.Header.Get("Authorization")
	// We should handle the error, someday
	tokenData, _ := utils.ReadJWT(token)
	comment.UserID = tokenData.User.ID

	// It's actually pretty stupid but there are quite a few
	// conversions we have to do in order to use the god damn
	// post id parameter, here we go

	// 1st we convert it to an uint64 which can't be used directly
	postID, err := strconv.ParseUint(ps.ByName("postID"), 10, 64)

	// and here we assign it converting it to a standard uint
	comment.PostID = uint(postID)

	// If the comment actually has content, the user exists and so does the post
	// then we can proceed to save it to the database
	// TBD: Lookup for the post in the db, also for the user to check for valid data
	if comment.Content != "" && comment.UserID != 0 && comment.PostID != 0 {
		db.Create(&comment)
	} else {
		utils.JSONMessage(w, "Content is required", http.StatusBadRequest)
		return
	}

	respBody, err := json.MarshalIndent(comment, "", " ")
	if err != nil {
		log.Fatalf("Error returning created comment: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusCreated)
}

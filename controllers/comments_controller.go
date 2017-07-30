package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
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
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Error reading token while creating post", http.StatusInternalServerError)
		return
	}

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
	// TBD: Check the db for valid user data

	// Find the post in the database
	var post models.Post
	db.First(&post, postID)

	// We make sure the post exists and that it's NOT locked
	if post.ID != 0 && post.IsLocked == false {
		if comment.Content != "" && comment.UserID != 0 && comment.PostID != 0 {
			db.Create(&comment)

			respBody, err := json.MarshalIndent(comment, "", " ")
			if err != nil {
				log.Fatalf("Error returning created comment: %v", err)
			}

			utils.JSONResponse(w, respBody, http.StatusCreated)
			return
		}

		utils.JSONMessage(w, "Content is required", http.StatusBadRequest)
		return
	}

	if post.IsLocked {
		utils.JSONMessage(w, "This post is locked and doesn't allow new comments", http.StatusForbidden)
		return
	}

	utils.JSONMessage(w, "The post was not found", http.StatusNotFound)
}

// DeleteComment destroys a comment with given ID
func DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// We need to verify that the one trying to delete the comment
	// is the same who created it

	// First we find the comment in the database
	var savedComment models.Comment
	db.First(&savedComment, ps.ByName("commentID"))

	// Then we get the user id from the JWT
	token := r.Header.Get("Authorization")
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Error reading token", http.StatusInternalServerError)
		return
	}

	userID := tokenData.User.ID

	// Compare the user id of who created the post
	// with the one stored in the token
	// then proceed with the delete
	if savedComment.UserID == userID {
		db.Delete(&savedComment)
		utils.JSONMessage(w, "Comment deleted", http.StatusOK)
		return
	}

	utils.JSONMessage(w, "An error ocurring deleting the comment", http.StatusNotFound)
}

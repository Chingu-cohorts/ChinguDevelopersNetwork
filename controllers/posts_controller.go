package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chingu-cohorts/ChinguCentral/models"
	"github.com/Chingu-cohorts/ChinguCentral/utils"
	"github.com/julienschmidt/httprouter"
)

// ListPosts returns a list of all the posts in the database
func ListPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var posts []models.Post
	db.Order("updated_at desc").Preload("User").Find(&posts)

	respBody, err := json.MarshalIndent(posts, "", " ")
	if err != nil {
		log.Fatalf("Error listing posts: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusOK)
}

// ShowPost returns the data for a post with given ID
func ShowPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var post models.Post
	db.First(&post, ps.ByName("id"))

	// If we found the post...
	if post.ID != 0 {
		respBody, err := json.MarshalIndent(post, "", " ")
		if err != nil {
			log.Fatalf("Error ocurred showing post: %v", err)
		}

		utils.JSONResponse(w, respBody, http.StatusOK)
		return
	}

	utils.JSONMessage(w, "Post not found", http.StatusNotFound)
}

// CreatePost saves a post in the database
func CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var post models.Post

	// I still have to get the user id from the token
	// and update it in the structure as well

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		utils.JSONMessage(w, "Wrong body", http.StatusBadRequest)
		return
	}

	if post.Title != "" && post.Content != "" && post.UserID != 0 {
		db.Create(&post)
	} else {
		utils.JSONMessage(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	respBody, err := json.MarshalIndent(post, "", " ")
	if err != nil {
		log.Fatalf("Error returning created post: %v", err)
	}

	utils.JSONResponse(w, respBody, http.StatusCreated)
}

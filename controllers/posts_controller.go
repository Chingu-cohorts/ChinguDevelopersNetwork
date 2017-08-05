package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/models"
	"github.com/Chingu-cohorts/ChinguDevelopersNetwork/utils"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

// ListPosts returns a list of all the posts in the database
func ListPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	var posts []models.Post
	db.Order("updated_at desc").Preload("User").Preload("Comments").Find(&posts)

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

	db.Preload("User").Preload("Comments.User").First(&post, ps.ByName("postID"))

	// If the post exists, it's ID must be different than 0
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

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		utils.JSONMessage(w, "Wrong body", http.StatusBadRequest)
		return
	}

	// Get the user id from the token and assign it
	token := r.Header.Get("Authorization")
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Error reading token while creating post", http.StatusInternalServerError)
		return
	}

	post.UserID = tokenData.User.ID

	// Check if there's a title and a content
	// It's done in db level as well, but it's a nice thing to do
	// at API level
	if post.Title != "" && post.Content != "" && post.UserID != 0 {
		// Generate a slug and assign it
		post.Slug = slug.Make(post.Title)

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

// UpdatePost performs an update to an existing post
func UpdatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// We have to get the user id from the token
	// Also get the post from the db and compare
	// the two ids
	var savedPost models.Post
	db.First(&savedPost, ps.ByName("postID"))

	token := r.Header.Get("Authorization")
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Error reading token while updating post", http.StatusInternalServerError)
		return
	}

	// Only if the one updating the post is the one who created it
	if savedPost.UserID == tokenData.User.ID {
		// The post that comes along with the request
		var post models.Post
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&post)
		if err != nil {
			utils.JSONMessage(w, "Invalid data", http.StatusBadRequest)
			return
		}

		db.Model(&savedPost).Updates(post)
		utils.JSONMessage(w, "Post updated", http.StatusOK)
		return
	}

	utils.JSONMessage(w, "You can't update other users posts", http.StatusUnauthorized)
}

// DeletePost performs a soft delete on a post
func DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := utils.InitDB()
	defer db.Close()

	// We want to do verify that the post user id matches the user id contained
	// in the JWT

	// First we find the post in the database
	var savedPost models.Post
	db.First(&savedPost, ps.ByName("postID"))

	// Then, we get the user object from the JWT
	token := r.Header.Get("Authorization")
	tokenData, err := utils.ReadJWT(token)
	if err != nil {
		utils.JSONMessage(w, "Error reading token", http.StatusInternalServerError)
		return
	}

	userID := tokenData.User.ID

	// Now we can compare the post user_id and the jwt user_id
	if savedPost.UserID == userID {
		db.Delete(&savedPost)
		utils.JSONMessage(w, "Post deleted", http.StatusOK)
		return
	}

	utils.JSONMessage(w, "An error ocurred deleting the post", http.StatusNotFound)
}

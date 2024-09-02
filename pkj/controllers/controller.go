package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Megidy/BloggingPlatformApi/pkj/models"
	"github.com/Megidy/BloggingPlatformApi/pkj/utils"
	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	Posts := models.GetAllPosts()
	response, err := json.Marshal(Posts)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
func GetPostById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	postId := vars["postId"]
	id, _ := strconv.ParseInt(postId, 0, 0)
	post, _ := models.GetPostById(id)
	response, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var NewPost models.Post
	utils.ParseBody(r, NewPost)
	post := NewPost.CreatePost()

	response, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	id, _ := strconv.ParseInt(postId, 0, 0)
	post := models.DeletePost(id)
	response, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var NewPost models.Post
	utils.ParseBody(r, NewPost)
	vars := mux.Vars(r)
	postId := vars["postId"]
	id, _ := strconv.ParseInt(postId, 0, 0)
	PostDetails, db := models.GetPostById(id)
	if NewPost.Title != "" {
		PostDetails.Title = NewPost.Title
	}
	if NewPost.Content != "" {
		PostDetails.Content = NewPost.Content
	}
	PostDetails.Category = NewPost.Category
	db.Save(&NewPost)
	response, err := json.Marshal(NewPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

package routes

import (
	"github.com/Megidy/BloggingPlatformApi/pkj/controllers"
	"github.com/gorilla/mux"
)

var InitRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts/", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{postId}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/posts/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{postId}", controllers.DeletePost).Methods("DELETE")

}

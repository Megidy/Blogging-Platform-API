package main

import (
	"log"
	"net/http"

	"github.com/Megidy/BloggingPlatformApi/pkj/config"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	config.Connect()
	log.Fatal(http.ListenAndServe(":8080", router))
}

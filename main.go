package main

import (
	"groupie-tracker/api"
	"groupie-tracker/handler"
	"log"
	"net/http"
)

func main() {
	var app handler.App
	var err error
	app.Artists, err = api.FetchAllData()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/static/", handler.StaticHandler)
	mux.HandleFunc("/", app.HomeHandler)
	mux.HandleFunc("/artist/", app.ArtistHandler)
	log.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

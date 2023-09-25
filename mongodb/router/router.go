package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"main.go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("routes")
	router.HandleFunc("/api/movies", controller.Getallmovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.Markaswatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.Deleteonemovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.Delteallmoveies).Methods("DELETE")

	return router
}

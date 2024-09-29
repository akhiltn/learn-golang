package router

import (
	"log"
	"net/http"

	"github.com/akhiltn/learn-golang/go-simple-rest-crud/api"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", api.ApiGetMovies).Methods(http.MethodGet)
	router.HandleFunc("/api/movie", api.ApiCreateMovie).Methods(http.MethodPost)
	router.HandleFunc("/api/movie/{id}", api.ApiMarkAsWatched).Methods(http.MethodPut)
	router.HandleFunc("/api/movie/{id}", api.ApiDeleteAMovie).Methods(http.MethodDelete)
	router.HandleFunc("/api/movies", api.ApiDeleteAllMovies).Methods(http.MethodDelete)
	log.Println("Router Initialized")
	return router
}

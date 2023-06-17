package router

import (
	"fmt"

	"github.com/djay-S/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Router for My Netflix APIs")

	router := mux.NewRouter()

	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie", controller.AddMovie).Methods("POST")
	router.HandleFunc("/movie/watched/{id}", controller.MarkMovieAsWatched).Methods("PUT")
	router.HandleFunc("/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}

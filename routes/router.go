package routes

import (
	"github.com/gorilla/mux"
	"github.com/manishSharma1-dev/goPractice/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/addonemovie", controllers.AddoneMovie).Methods("POST")
	r.HandleFunc("/api/getallmovies", controllers.GetallMovies).Methods("GET")
	r.HandleFunc("/api/updateonemovie/{id}", controllers.UpdateoneMovietomarkwatch).Methods("PUT")
	r.HandleFunc("/api/deleteonemovie/{id}", controllers.DeleteOnemovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovies", controllers.DeleteAllMovies).Methods("DELETE")

	return r

}

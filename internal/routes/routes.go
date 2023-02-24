package routes

import (
	"github.com/gorilla/mux"
	"github.com/rogaliiik/playlist/internal/controllers"
)

var RegisterRoutes = func(router *mux.Router) {

	router.HandleFunc("/song", controllers.GetAllSongs).Methods("GET")
	router.HandleFunc("/song/{songId}", controllers.GetSongById).Methods("GET")
	router.HandleFunc("/song", controllers.CreateSong).Methods("POST")
	router.HandleFunc("/song/{songId}", controllers.UpdateSong).Methods("PUT")
	router.HandleFunc("/song/{songId}", controllers.DeleteSong).Methods("DELETE")

	router.HandleFunc("/play", controllers.PlaySong).Methods("GET")
	router.HandleFunc("/pause", controllers.PauseSong).Methods("GET")
	router.HandleFunc("/next", controllers.NextSong).Methods("GET")
	router.HandleFunc("/prev", controllers.PrevSong).Methods("GET")

}

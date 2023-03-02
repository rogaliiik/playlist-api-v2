package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/rogaliiik/playlist/internal/playlist"
	"github.com/rogaliiik/playlist/internal/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)

	internal.APIServer.Playlist.Wg.Add(1)
	defer internal.APIServer.Playlist.Wg.Wait()

	go internal.APIServer.Playlist.Broadcast()

	log.Println("Server is working at localhost:8080")
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}

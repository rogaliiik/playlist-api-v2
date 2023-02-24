package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rogaliiik/playlist/internal/playlist"
	"github.com/rogaliiik/playlist/internal/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)

	p := internal.NewPlaylist()
	p.Wg.Add(1)
	defer p.Wg.Wait()

	p.AddSong(internal.NewSong("Song 1", 1))
	p.AddSong(internal.NewSong("Song 2", 1))
	p.AddSong(internal.NewSong("Song 3", 1))

	go p.Broadcast()

	fmt.Println("Server is working at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

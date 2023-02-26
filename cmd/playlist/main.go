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

	p := internal.GetPlaylist()
	p.Wg.Add(1)
	defer p.Wg.Wait()

	go p.Broadcast()

	fmt.Println("Server is working at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

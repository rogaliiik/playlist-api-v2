package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rogaliiik/playlist/internal/models"
	internal "github.com/rogaliiik/playlist/internal/playlist"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rogaliiik/playlist/utility"
)

var (
	playlistDB *models.PlaylistDB
)

func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	Songs := internal.GetAllSongs()

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(Songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func GetSongById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songId := vars["songId"]
	id, err := strconv.Atoi(songId)
	if err != nil {
		log.Println(err)
		return
	}
	song := internal.GetSongById(uint(id))

	res, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSong(w http.ResponseWriter, r *http.Request) {
	newSong := &models.Song{}
	utility.ParseBody(r, newSong)

	song := internal.AddSong(newSong)
	res, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songId := vars["songId"]
	id, err := strconv.Atoi(songId)
	if err != nil {
		log.Println(err)
		return
	}
	deleted, err := internal.DeleteSong(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(deleted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	var updateSong = &models.Song{}
	utility.ParseBody(r, updateSong)

	vars := mux.Vars(r)
	songId := vars["songId"]
	id, err := strconv.Atoi(songId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	song := internal.UpdateSong(id, updateSong)
	res, err := json.Marshal(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PlaySong(w http.ResponseWriter, r *http.Request) {
	playlistDB = internal.PlaySong()
	res, err := json.Marshal(playlistDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PauseSong(w http.ResponseWriter, r *http.Request) {
	playlistDB = internal.PauseSong()
	res, err := json.Marshal(playlistDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func NextSong(w http.ResponseWriter, r *http.Request) {
	next := internal.NextSong()

	res, err := json.Marshal(next)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PrevSong(w http.ResponseWriter, r *http.Request) {
	prev := internal.PrevSong()

	res, err := json.Marshal(prev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Status(w http.ResponseWriter, r *http.Request) {
	playlistDB = internal.GetPlaylistDB()
	res, err := json.Marshal(playlistDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func init() {
	playlistDB = internal.GetPlaylistDB()
}

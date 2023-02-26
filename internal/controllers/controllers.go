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
	playlistDB *models.Playlist
	playlist   *internal.Player
)

func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	Songs := models.GetAllSongs()

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
	song := models.GetSongById(uint(id))

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

	song := playlistDB.AddSong(newSong)
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
	deleted, err := models.DeleteSong(uint(id))
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
	song := models.UpdateSong(id, updateSong)
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
	playlist.Play()
	playlistDB.Play()
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
	playlist.Pause()
	playlistDB.Pause()
	playlistDB.State = internal.Pause
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
	next := playlistDB.NextSong()

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
	prev := playlistDB.PrevSong()

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
	playlistDB = models.GetPlaylist()
	playlist = internal.GetPlaylist()
}

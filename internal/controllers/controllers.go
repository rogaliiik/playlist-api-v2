package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rogaliiik/playlist/internal/playlist"
	"github.com/rogaliiik/playlist/utility"
)

var p *internal.Playlist

func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	var Songs []map[string]string
	s := p.Head
	for s != nil {
		Songs = append(Songs, s.FieldsToMap())
		s = s.Next
	}
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
	song := p.Head
	ID, err := strconv.Atoi(songId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	for song != nil {
		if song.ID.ID() == uint32(ID) {
			break
		}
		song = song.Next
	}
	res, err := json.Marshal(song.FieldsToMap())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSong(w http.ResponseWriter, r *http.Request) {
	newSong := internal.NewSong("", 0)
	utility.ParseBody(r, newSong)

	p.AddSong(newSong)
	res, err := json.Marshal(newSong.FieldsToMap())
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
	ID, err := strconv.Atoi(songId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deleted := p.RemoveSong(uint32(ID))

	if deleted == nil {
		return
	}
	res, err := json.Marshal(deleted.FieldsToMap())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	var updateSong = internal.NewSong("", 0)
	utility.ParseBody(r, updateSong)

	vars := mux.Vars(r)
	songId := vars["songId"]
	song := p.Head
	ID, err := strconv.Atoi(songId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	for song != nil {
		if song.ID.ID() == uint32(ID) {
			break
		}
		song = song.Next
	}
	if song == nil {
		return
	}
	if updateSong.Title != "" {
		song.Title = updateSong.Title
	}
	if updateSong.Duration != 0 {
		song.Duration = updateSong.Duration
	}

	res, err := json.Marshal(song.FieldsToMap())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PlaySong(w http.ResponseWriter, r *http.Request) {
	p.Play()
	res, err := json.Marshal(map[string]string{"state": fmt.Sprintf("%d", p.State)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PauseSong(w http.ResponseWriter, r *http.Request) {
	p.Pause()
	res, err := json.Marshal(map[string]string{"state": fmt.Sprintf("%d", p.State)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func NextSong(w http.ResponseWriter, r *http.Request) {
	p.Next()
	if p.Current == nil {
		return
	}
	res, err := json.Marshal(p.Current.FieldsToMap())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PrevSong(w http.ResponseWriter, r *http.Request) {
	p.Prev()
	if p.Current == nil {
		return
	}
	res, err := json.Marshal(p.Current.FieldsToMap())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func init() {
	p = internal.GetPlaylist()
}

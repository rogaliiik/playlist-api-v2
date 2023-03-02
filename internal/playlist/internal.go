package internal

import (
	"fmt"
	"sync"
	"time"

	"github.com/rogaliiik/playlist/internal/config"
	"github.com/rogaliiik/playlist/internal/models"
)

var APIServer *Server

type Server struct {
	Store    *config.Store
	Playlist *Playlist
}

func NewServer(store *config.Store, playlist *Playlist) *Server {
	return &Server{store, playlist}
}

const (
	Pause = 0
	Play  = 1
)

type Playlist struct {
	playlistDB *models.PlaylistDB
	Wg         sync.WaitGroup
}

func (p *Playlist) Shutdown() {
	fmt.Println("Exit...")
	p.Wg.Done()
}

func (p *Playlist) Broadcast() {
	tick := time.Tick(1 * time.Second)
	for {
		p.playlistDB = GetPlaylistDB()
		song := GetSongById(p.playlistDB.Current)
		for p.playlistDB.Current != 0 && p.playlistDB.Timecode < song.Duration {
			select {
			case <-tick:
				p.playlistDB = GetPlaylistDB()
				song = GetSongById(p.playlistDB.Current)
				if p.playlistDB.State == Play {
					fmt.Println("id:", song.ID, "name:", song.Title, "timecode:", p.playlistDB.Timecode)
					p.playlistDB.Timecode += 1
					APIServer.Store.DB.Save(&p.playlistDB)
				}
			}
		}
		if p.playlistDB.State == Play {
			NextSong()
			p.playlistDB = GetPlaylistDB()
		}
	}
}

func init() {
	db := config.Connect()
	APIServer = NewServer(db, &Playlist{})
	APIServer.Store.DB.AutoMigrate(&models.Song{}, &models.PlaylistDB{})
}

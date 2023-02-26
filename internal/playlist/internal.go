package internal

import (
	"fmt"
	"github.com/rogaliiik/playlist/internal/models"
	"sync"
	"time"
)

var (
	playlist *Playlist
)

// TODO: tests, Docker, errors, comments
const (
	Pause = 0
	Play  = 1
)

type Playlist struct {
	playlistDB *models.PlaylistDB
	Wg         sync.WaitGroup
}

func GetPlaylist() *Playlist {
	return playlist
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
					models.Store.DB.Save(&p.playlistDB)
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
	playlist = &Playlist{}
}

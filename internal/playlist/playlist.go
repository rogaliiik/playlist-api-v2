package internal

import (
	"fmt"
	"github.com/rogaliiik/playlist/internal/models"
	"sync"
	"time"
)

var (
	playlistDB *models.Playlist
	player     *Player
)

const (
	Pause = 0
	Play  = 1
)

type Player struct {
	State    uint
	Current  uint
	timecode uint
	Wg       sync.WaitGroup
}

func GetPlaylist() *Player {
	return player
}

func (p *Player) Play() {
	p.State = Play
}

func (p *Player) Pause() {
	p.State = Pause
}

func (p *Player) Next() {
	p.timecode = 0
	p.Current = playlistDB.Current
}

func (p *Player) Prev() {
	p.timecode = 0
	p.Current = playlistDB.Current
}

func (p *Player) Shutdown() {
	fmt.Println("Exit...")
	p.Wg.Done()
}

func (p *Player) Broadcast() {
	for {
		song := models.GetSongById(p.Current)
		tick := time.Tick(1 * time.Second)
		//fmt.Println("p.Current", p.Current, p.timecode)
		for p.Current != 0 && p.timecode <= song.Duration {
			select {
			case <-tick:
				if p.State == Play {
					fmt.Println(song.ID, song.Title, p.timecode)
					p.timecode += 1
					playlistDB.Update(playlistDB.State, playlistDB.Current, playlistDB.Timecode+1)
				}
			}
		}
		if p.State == Play {
			p.Next()
		}
	}
}

func init() {
	playlistDB = models.GetPlaylist()
	player = &Player{
		State:    playlistDB.State,
		Current:  playlistDB.Current,
		timecode: playlistDB.Timecode,
	}
}

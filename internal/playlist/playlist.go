package internal

import (
	"fmt"
	"sync"
	"time"
)

const (
	Pause = 1
	Play  = 0
)

type Playlist struct {
	state       int
	currentSong *Song
	lastSong    *Song
	timecode    int64
	PlayChan    chan int64
	PauseChan   chan int64
	mutex       sync.Mutex
}

func NewPlaylist() *Playlist {
	firstSong := &Song{}
	return &Playlist{
		state:       Pause,
		currentSong: firstSong,
		lastSong:    firstSong,
		PlayChan:    make(chan int64),
		PauseChan:   make(chan int64),
		timecode:    0,
	}
}

type Song struct {
	Title    string
	Duration int64
	Prev     *Song
	Next     *Song
}

func NewSong(title string, duration int64) *Song {
	return &Song{
		Title:    title,
		Duration: duration,
	}
}

func (p *Playlist) AddSong(song *Song) {
	p.lastSong.Next = song
	song.Prev = p.lastSong
	p.lastSong = song
}

func (p *Playlist) Next() {
	p.timecode = 0
	if p.currentSong.Next != nil {
		p.currentSong = p.currentSong.Next
		go p.Play()
		return
	}
	fmt.Println("there are no songs after this one")
	p.PlayChan <- 1
	return
}

func (p *Playlist) Prev() {
	p.timecode = 0
	if p.currentSong.Prev.Title != "" {
		p.currentSong = p.currentSong.Prev
		go p.Play()
		return
	}
	fmt.Println("there are no songs before this one")
	p.PlayChan <- 1
	return
}

//TODO: add mutex correctly
func (p *Playlist) Play() {
	if p.currentSong.Title == "" {
		if p.currentSong.Next.Title == "" {
			p.PlayChan <- Pause
			fmt.Println("there are no songs")
			return
		}
		p.currentSong = p.currentSong.Next
	}
	p.state = Play
	fmt.Println(p.currentSong.Title, "is playing")

	for p.timecode <= p.currentSong.Duration {
		select {
		case <-p.PauseChan:
			p.PlayChan <- Pause
			return
		default:
			fmt.Println(p.timecode, p.currentSong.Duration)
			time.Sleep(time.Second)
			p.timecode += 1
		}
	}
	p.Next()
	return
}

func (p *Playlist) Pause() {
	if p.state == Play {
		p.state = Pause
		p.PauseChan <- Pause
	}
	fmt.Println(p.currentSong.Title, "paused")
}

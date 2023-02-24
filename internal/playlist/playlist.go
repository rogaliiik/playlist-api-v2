package internal

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

type State int8

var p Playlist

const (
	Pause State = 0
	Play  State = 1
)

type Song struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Duration int       `json:"duration"`
	Prev     *Song
	Next     *Song
}

func NewSong(title string, duration int) *Song {
	return &Song{Title: title, Duration: duration, ID: uuid.New()}
}

func (s *Song) FieldsToMap() map[string]string {
	return map[string]string{
		"id":       strconv.Itoa(int(s.ID.ID())),
		"title":    s.Title,
		"duration": strconv.Itoa(s.Duration)}
}

type Playlist struct {
	State    State
	Current  *Song
	Tail     *Song
	Head     *Song
	timecode int
	mutex    sync.Mutex
	Wg       sync.WaitGroup
}

func NewPlaylist() *Playlist {
	p = Playlist{State: Pause, Tail: nil, Head: nil, Current: nil}
	return &p
}

func GetPlaylist() *Playlist {
	return &p
}

func (p *Playlist) AddSong(s *Song) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	fmt.Println("New song:", s.Title)
	node := s
	if p.Tail == nil {
		p.Tail = node
		p.Current = node
		p.Head = node
	} else {
		p.Tail.Next = node
		node.Prev = p.Tail
		p.Tail = node
	}
}

func (p *Playlist) RemoveSong(id uint32) *Song {
	song := p.Head
	for song != nil {
		if song.ID.ID() == id {
			if p.Current.ID.ID() == id {
				return nil
			}
			prev := song.Prev
			next := song.Next
			if prev != nil {
				prev.Next = next
			}
			if next != nil {
				next.Prev = prev
			}
			return song
		}
		song = song.Next
	}
	return nil
}

func (p *Playlist) Play() {
	p.State = Play

}

func (p *Playlist) Pause() {
	p.State = Pause

}

func (p *Playlist) Next() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.timecode = 0
	if p.Current != nil {
		if p.Current.Next == nil {
			p.Current = p.Head
		} else {
			p.Current = p.Current.Next
		}
		fmt.Println("Next:", p.Current.Title)
		p.Play()
		return nil
	}
	return fmt.Errorf("there are no songs in playlist")
}

func (p *Playlist) Prev() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.timecode = 0
	if p.Current != nil {
		if p.Current.Prev == nil {
			p.Current = p.Tail
		} else {
			p.Current = p.Current.Prev
		}
		fmt.Println("Previous:", p.Current.Title)
		p.Play()
		return nil
	}
	return fmt.Errorf("there are no songs in playlist")
}

func (p *Playlist) Shutdown() {
	fmt.Println("Exit...")
	p.Wg.Done()
}

func (p *Playlist) Broadcast() {
	for {
		tick := time.Tick(1 * time.Second)
		for p.Current != nil && p.timecode <= p.Current.Duration {
			select {
			case <-tick:
				if p.State == Play {
					fmt.Println(p.Current.ID.ID(), p.Current.Title, p.timecode)
					p.timecode += 1
				}
			}
		}
		if p.State == Play {
			p.Next()
		}
	}
}

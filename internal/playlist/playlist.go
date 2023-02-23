package internal

import (
	"fmt"
	"sync"
	"time"
)

type State int8

const (
	Pause = 0
	Play  = 1
)

type Song struct {
	title    string
	duration int64
	prev     *Song
	next     *Song
}

func NewSong(title string, duration int64) *Song {
	return &Song{title: title, duration: duration}
}

type Playlist struct {
	state    State
	current  *Song
	tail     *Song
	head     *Song
	timecode int64
	mutex    sync.Mutex
	Wg       sync.WaitGroup
}

func NewPlaylist() *Playlist {
	return &Playlist{state: Pause, tail: nil, head: nil, current: nil}
}

func (p *Playlist) AddSong(s *Song) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	fmt.Println("New song:", s.title)
	node := s
	if p.tail == nil {
		p.tail = node
		p.current = node
		p.head = node
	} else {
		p.tail.next = node
		node.prev = p.tail
		p.tail = node
	}
}

func (p *Playlist) Play() {
	p.state = Play

}

func (p *Playlist) Pause() {
	p.state = Pause

}

func (p *Playlist) Next() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.timecode = 0
	if p.current != nil {
		if p.current.next == nil {
			p.current = p.head
		} else {
			p.current = p.current.next
		}
		fmt.Println("Next:", p.current.title)
		p.Play()
		return nil
	}
	return fmt.Errorf("there are no songs in playlist")
}

func (p *Playlist) Prev() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.timecode = 0
	if p.current != nil {
		if p.current.prev == nil {
			p.current = p.tail
		} else {
			p.current = p.current.prev
		}
		fmt.Println("Previous:", p.current.title)
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
		for p.current != nil && p.timecode <= p.current.duration {
			select {
			case <-tick:
				if p.state == Play {
					fmt.Println(p.current.title, p.timecode)
					p.timecode += 1
				}
			}
		}
		if p.state == Play {
			p.Next()
		}
	}
}

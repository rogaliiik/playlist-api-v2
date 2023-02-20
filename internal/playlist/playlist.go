package internal

import (
	"fmt"
	"sync"
	"time"
)

//TODO: add logs

type State int8

const (
	Pause = 0
	Play  = 1
)

type Node struct {
	song *Song
	prev *Node
	next *Node
}

type Song struct {
	title    string
	duration int64
}

func NewSong(title string, duration int64) *Song {
	return &Song{title: title, duration: duration}
}

type Playlist struct {
	state    State
	current  *Node
	tail     *Node
	head     *Node
	timecode int64
	mu       sync.Mutex
	Wg       sync.WaitGroup
}

func NewPlaylist() *Playlist {
	return &Playlist{state: Pause, tail: nil, head: nil, current: nil}
}

func (p *Playlist) Play() {
	p.state = Play

}

func (p *Playlist) Pause() {
	p.state = Pause

}

func (p *Playlist) Next() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.timecode = 0
	if p.current != nil {
		if p.current.next == nil {
			p.current = p.head
		} else {
			p.current = p.current.next
		}
		fmt.Println("Next:", p.current.song.title)
		p.Play()
	}
}

func (p *Playlist) Prev() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.timecode = 0
	if p.current != nil {
		if p.current.prev == nil {
			p.current = p.tail
		} else {
			p.current = p.current.prev
		}
		fmt.Println("Previous:", p.current.song.title)
		p.Play()
	}
}

func (p *Playlist) AddSong(s *Song) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Println("New song:", s.title)
	node := &Node{song: s}
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

func (p *Playlist) Shutdown() {
	fmt.Println("Exit...")
	p.Wg.Done()
}

func (p *Playlist) Broadcast() {
	for {
		for p.current != nil && p.timecode <= p.current.song.duration {
			if p.state == Play {
				fmt.Println(p.current.song.title, p.timecode)
				time.Sleep(time.Second)
				p.timecode += 1
			}
		}
		if p.state == Play {
			p.Next()
		}
	}
}

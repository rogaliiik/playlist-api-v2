package main

import (
	internal "github.com/rogaliiik/playlist/internal/playlist"
)

func main() {
	p := internal.NewPlaylist()
	p.Wg.Add(1)
	defer p.Wg.Wait()

	p.AddSong(internal.NewSong("Song 1", 1))
	p.AddSong(internal.NewSong("Song 2", 1))
	p.AddSong(internal.NewSong("Song 3", 1))

	go p.Broadcast()
	p.Next()
	p.Prev()
	p.AddSong(internal.NewSong("Song 4", 1))
	p.Next()

}

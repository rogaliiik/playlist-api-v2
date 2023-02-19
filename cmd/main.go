package main

import (
	internal "github.com/rogaliiik/playlist/internal/playlist"
	"time"
)

func main() {
	p := internal.NewPlaylist()

	p.AddSong(&internal.Song{
		Title:    "Song 1",
		Duration: 2,
	})
	p.AddSong(&internal.Song{
		Title:    "Song 2",
		Duration: 2,
	})
	p.AddSong(&internal.Song{
		Title:    "Song 3",
		Duration: 2,
	})

	go p.Play()
	time.Sleep(1)
	p.Pause()
	<-p.PlayChan

	p.Next()
	p.Next()
	<-p.PlayChan

}

package models

import (
	"github.com/rogaliiik/playlist/internal/config"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Song struct {
	gorm.Model
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Next     int    `json:"next"`
	Prev     int    `json:"prev"`
}

type Playlist struct {
	gorm.Model
	State    int `json:"state"`
	Current  int `json:"current"`
	Timecode int `json:"timecode"`
}

func GetPlaylist() *Playlist {
	var p *Playlist
	db.First(&p)
	return p
}

func GetCurrentSong(id int) *Song {
	var song *Song
	db.Where("ID=?", id).Find(&song)
	return song
}

func GetHeadSong() *Song {
	var head *Song
	db.First(&head)
	return head
}

func GetTailSong() *Song {
	var tail *Song
	db.Last(&tail)
	return tail
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Song{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Playlist{})
	if err != nil {
		log.Println(err)
	}
}

package models

import (
	"errors"
	"fmt"
	"github.com/rogaliiik/playlist/internal/config"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Playlist struct {
	gorm.Model
	State    uint `json:"state" gorm:"default:0"`
	Current  uint `json:"current" gorm:"default:0"`
	Timecode uint `json:"timecode" gorm:"default:0"`
}

func GetPlaylist() *Playlist {
	var p *Playlist
	err := db.First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&Playlist{})
		db.First(&p)
		return p
	}
	return p
}

func (p *Playlist) Update(state, current, timecode uint) {
	p.State = state
	p.Current = current
	p.Timecode = timecode
	db.Save(&p)
}

func (p *Playlist) AddSong(song *Song) *Song {
	var tail *Song
	err := db.Last(&tail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newSong := CreateSong(song)
		p.Current = newSong.ID
		db.Save(&p)
		return newSong
	}
	fmt.Println(tail.ID, tail.Next, tail.Prev)
	newSong := CreateSong(song)
	tail.Next = newSong.ID
	newSong.Prev = tail.ID
	db.Save(&tail)
	db.Save(&newSong)
	return newSong
}

func (p *Playlist) NextSong() *Song {
	if p.Current != 0 {
		cur := GetSongById(p.Current)
		if cur.Next != 0 {
			song := GetSongById(cur.Next)
			p.Update(p.State, song.ID, 0)
			return song
		}
		var head *Song
		db.First(&head)
		p.Update(p.State, head.ID, 0)
		return head
	}
	return nil
}

func (p *Playlist) PrevSong() *Song {
	if p.Current != 0 {
		cur := GetSongById(p.Current)
		if cur.Prev != 0 {
			song := GetSongById(cur.Prev)
			p.Update(p.State, song.ID, 0)
			return song
		}
		var tail *Song
		db.First(&tail)
		p.Update(p.State, tail.ID, 0)
		return tail
	}
	return nil
}

func (p *Playlist) Play() {
	p.Update(1, p.Current, p.Timecode)
}

func (p *Playlist) Pause() {
	p.Update(0, p.Current, p.Timecode)
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Song{}, &Playlist{})
	if err != nil {
		log.Println(err)
	}
}

package internal

import (
	"errors"
	"github.com/rogaliiik/playlist/internal/models"
	"gorm.io/gorm"
	"log"
)

func GetPlaylistDB() *models.PlaylistDB {
	var p *models.PlaylistDB

	err := models.Store.DB.First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		models.Store.DB.Create(&models.PlaylistDB{})
		models.Store.DB.First(&p)
		return p
	}
	return p
}

func AddSong(song *models.Song) *models.Song {
	p := GetPlaylistDB()
	var tail *models.Song
	err := models.Store.DB.Last(&tail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newSong := CreateSong(song)
		p.Current = newSong.ID
		models.Store.DB.Save(&p)
		return newSong
	}
	newSong := CreateSong(song)
	tail.Next = newSong.ID
	newSong.Prev = tail.ID
	models.Store.DB.Save(&tail)
	models.Store.DB.Save(&newSong)
	return newSong
}

func NextSong() *models.Song {
	p := GetPlaylistDB()
	if p.Current != 0 {
		cur := GetSongById(p.Current)
		if cur.Next != 0 {
			song := GetSongById(cur.Next)
			p.Current = song.ID
			p.Timecode = 0
			models.Store.DB.Save(&p)
			return song
		}
		var head *models.Song
		models.Store.DB.First(&head)
		p.Current = head.ID
		p.Timecode = 0
		models.Store.DB.Save(&p)
		return head
	}
	return nil
}

func PrevSong() *models.Song {
	p := GetPlaylistDB()
	if p.Current != 0 {
		cur := GetSongById(p.Current)
		if cur.Prev != 0 {
			song := GetSongById(cur.Prev)
			p.Current = song.ID
			p.Timecode = 0
			models.Store.DB.Save(&p)
			return song
		}
		var tail *models.Song
		models.Store.DB.Last(&tail)
		p.Current = tail.ID
		p.Timecode = 0
		models.Store.DB.Save(&p)
		return tail
	}
	return nil
}

func PlaySong() *models.PlaylistDB {
	p := GetPlaylistDB()
	p.State = Play
	models.Store.DB.Save(&p)
	log.Println("playlist is playing")
	return p
}

func PauseSong() *models.PlaylistDB {
	p := GetPlaylistDB()
	p.State = Pause
	models.Store.DB.Save(&p)
	log.Println("playlist is paused")
	return p
}

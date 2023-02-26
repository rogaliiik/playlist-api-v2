package models

import (
	"gorm.io/gorm"
	"log"
)

type Song struct {
	gorm.Model
	Title    string `json:"title"`
	Duration uint   `json:"duration"`
	Next     uint   `json:"next" gorm:"default:0"`
	Prev     uint   `json:"prev" gorm:"default:0"`
}

func CreateSong(s *Song) *Song {
	db.Create(&s)
	return s
}

func GetAllSongs() []Song {
	var Songs []Song
	db.Find(&Songs)
	return Songs
}

func GetSongById(id uint) *Song {
	var getSong Song
	db.Where("ID=?", id).Find(&getSong)
	return &getSong
}

func UpdateSong(id int, updateSong *Song) Song {
	song := GetSongById(uint(id))

	if updateSong.Title != "" {
		song.Title = updateSong.Title
	}
	if updateSong.Duration != 0 {
		song.Duration = updateSong.Duration
	}
	db.Save(&song)
	return *song
}

func DeleteSong(id uint) (*Song, error) {
	song := GetSongById(id)
	err := db.Delete(&Song{}, id).Error
	if err != nil {
		log.Println(err)
		return &Song{}, err
	}

	prev := GetSongById(song.Prev)
	next := GetSongById(song.Next)
	prev.Next = song.Next
	next.Prev = song.Prev
	db.Save(&prev)
	db.Save(&next)
	return song, nil
}

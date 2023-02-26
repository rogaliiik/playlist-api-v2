package internal

import (
	"github.com/rogaliiik/playlist/internal/models"
	"log"
)

func CreateSong(s *models.Song) *models.Song {
	models.Store.DB.Create(&s)
	return s
}

func GetAllSongs() []models.Song {
	var Songs []models.Song
	models.Store.DB.Find(&Songs)
	return Songs
}

func GetSongById(id uint) *models.Song {
	var getSong models.Song
	models.Store.DB.Where("ID=?", id).Find(&getSong)
	return &getSong
}

func UpdateSong(id int, updateSong *models.Song) models.Song {
	song := GetSongById(uint(id))

	if updateSong.Title != "" {
		song.Title = updateSong.Title
	}
	if updateSong.Duration != 0 {
		song.Duration = updateSong.Duration
	}
	models.Store.DB.Save(&song)
	return *song
}

func DeleteSong(id uint) (*models.Song, error) {
	song := GetSongById(id)
	if song.ID == GetPlaylistDB().Current {
		log.Println("impossible to delete current song")
		return &models.Song{}, nil
	}
	err := models.Store.DB.Delete(&models.Song{}, id).Error
	if err != nil {
		log.Println(err)
		return &models.Song{}, err
	}

	prev := GetSongById(song.Prev)
	next := GetSongById(song.Next)
	prev.Next = song.Next
	next.Prev = song.Prev
	models.Store.DB.Save(&prev)
	models.Store.DB.Save(&next)
	return song, nil
}

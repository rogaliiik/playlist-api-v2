package models

import (
	"github.com/rogaliiik/playlist/internal/config"
	"gorm.io/gorm"
	"log"
)

var Store config.Store

type Song struct {
	gorm.Model
	Title    string `json:"title"`
	Duration uint   `json:"duration"`
	Next     uint   `json:"next" gorm:"default:0"`
	Prev     uint   `json:"prev" gorm:"default:0"`
}

func init() {
	config.Connect()
	Store = config.GetDB()
	err := Store.DB.AutoMigrate(&Song{}, &PlaylistDB{})
	if err != nil {
		log.Println(err)
	}
}

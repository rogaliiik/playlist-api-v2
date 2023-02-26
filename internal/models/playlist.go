package models

import (
	"gorm.io/gorm"
)

type PlaylistDB struct {
	gorm.Model
	State    uint `json:"state" gorm:"default:0"`
	Current  uint `json:"current" gorm:"default:0"`
	Timecode uint `json:"timecode" gorm:"default:0"`
}

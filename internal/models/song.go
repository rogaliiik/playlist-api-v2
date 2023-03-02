package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title    string `json:"title"`
	Duration uint   `json:"duration"`
	Next     uint   `json:"next" gorm:"default:0"`
	Prev     uint   `json:"prev" gorm:"default:0"`
}

package models

import "time"

type FriendLink struct {
	MODEL
	Name        string    `gorm:"size:50" json:"name"`
	Description string    `gorm:"size:50" json:"description"`
	Logo        string    `gorm:"size:200" json:"logo"`
	Url         string    `gorm:"size:200" json:"url"`
	IsTop       bool      `json:"is_top"`
	TopTime     time.Time `json:"top_time"`
}

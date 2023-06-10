package models

type Announcement struct {
	MODEL
	Content string `gorm:"mediumtext" json:"content"`
}

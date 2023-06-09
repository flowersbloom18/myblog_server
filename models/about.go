package models

type About struct {
	MODEL
	Content string `gorm:"mediumtext" json:"content"`
}

package models

type Music struct {
	MODEL
	Name   string `gorm:"size:50" json:"name"`
	Author string `gorm:"size:50" json:"author"`
	Url    string `gorm:"size:200" json:"url"`
	Cover  string `gorm:"size:200" json:"cover"`
	Status bool   `json:"status"`
	Sort   int    `gorm:"type:int(6)" json:"sort"`
}

package models

type Info struct {
	MODEL
	Content string `gorm:"type:mediumtext" json:"content"`
	TypeId  int    `gorm:"size:10" json:"type_id"`
}

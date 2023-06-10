package models

type Info struct {
	MODEL
	Content string `gorm:"type:mediumtext" json:"content"`
	TypeId  int    `gorm:"type:int(6)" json:"type_id"`
}

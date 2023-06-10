package models

// （ID，createTime，updateTime，blogID，userID）
type Collect struct {
	MODEL
	BlogID int `gorm:"type:int(6)" json:"blog_id"`
	UserID int `gorm:"type:int(6)" json:"user_id"`
}

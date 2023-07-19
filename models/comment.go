package models

import "myblog_server/models/model_type"

type Comment struct {
	MODEL
	Content   string `gorm:"type:mediumtext" json:"content"` // 评论内容
	UserID    uint   `gorm:"type:int(6)" json:"user_id"`     // 评论用户ID
	IPAddress string `gorm:"size:50" json:"ip_address"`      // IP属地【评论时刻的信息，不会发生改变】

	PageType model_type.PageType `gorm:"type:int(6)" json:"page_type"` // 评论页面的类型
	Page     string              `gorm:"size:200" json:"page"`         // 评论页面
	IsAdmin  bool                `json:"is_admin"`                     // 是否为管理员
	FatherID uint                `gorm:"type:int(6)" json:"father_id"` // 父级ID
	PanelID  uint                `gorm:"type:int(6)" json:"panel_id"`  // 面板ID
}

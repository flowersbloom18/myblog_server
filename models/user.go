package models

import (
	"myblog_server/models/model_type"
	"time"
)

type User struct {
	//MODEL

	// 展示时间
	ID        uint      `gorm:"primarykey" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"`           // 创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                                 // 更新时间

	// 用户名添加唯一索引,防止重复
	UserName       string                `gorm:"uniqueIndex;size:36" json:"user_name,select(info)"` // 用户名
	Password       string                `gorm:"size:128" json:"-"`                                 // 密码
	NickName       string                `gorm:"size:36" json:"nick_name,select(c|info)"`           // 昵称
	Email          string                `gorm:"size:36" json:"email,select(c|info)"`               // 邮箱
	Avatar         string                `gorm:"size:200" json:"avatar,select(c|info)"`             // 头像
	Role           model_type.Role       `gorm:"type:int(6)" json:"role,select(info)"`              // 权限
	RegisterOrigin model_type.SignStatus `gorm:"type:int(6)" json:"register_origin,select(info)"`   // 注册来源
	IP             string                `gorm:"size:36" json:"ip,select(info)"`                    // IP地址
	Address        string                `gorm:"size:36" json:"address,select(c|info)"`             // 地址
	Device         string                `gorm:"size:36" json:"device,select(c|info)"`              // 登录设备
}

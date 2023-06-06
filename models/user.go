package models

import "myblog_server/models/model_type"

type User struct {
	MODEL
	UserName       string                `gorm:"size:36" json:"user_name"`                    // 用户名
	Password       string                `gorm:"size:128" json:"-"`                           // 密码
	NickName       string                `gorm:"size:36" json:"nick_name,select(c|info)"`     // 昵称
	Email          string                `gorm:"size:36" json:"email,select(c|info)"`         // 邮箱
	Avatar         string                `gorm:"size:200" json:"avatar,select(c|info)"`       // 头像
	Role           model_type.Role       `gorm:"size:36" json:"role,select(info)"`            // 权限
	RegisterOrigin model_type.SignStatus `gorm:"size:36" json:"register_origin,select(info)"` // 注册来源
	IP             string                `gorm:"size:36" json:"ip"`                           // IP地址
	Address        string                `gorm:"size:36" json:"address,select(c|info)"`       // 地址
	Device         string                `gorm:"size:36" json:"device,select(c|info)"`        // 登录设备
}

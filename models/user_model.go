package models

import "myblog_server/models/model_type"

type UserModel struct {
	MODEL
	UserName       string                `gorm:"size:36" json:"user_name"`       //用户名
	NickName       string                `gorm:"size:36" json:"nick_name"`       //昵称
	Password       string                `gorm:"size:128" json:"-"`              //密码
	Email          string                `gorm:"size:36" json:"email"`           //邮箱
	Avatar         string                `gorm:"size:36" json:"avatar"`          //头像
	Role           model_type.Role       `gorm:"size:36" json:"role"`            //权限
	RegisterOrigin model_type.SignStatus `gorm:"size:36" json:"register_origin"` //注册来源
	IP             string                `gorm:"size:36" json:"ip"`              //IP地址
	Address        string                `gorm:"size:36" json:"address"`         //地址
}

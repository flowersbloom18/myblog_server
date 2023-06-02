package models

import (
	"myblog_server/models/model_type"
)

type LogModel struct {
	MODEL
	UserName  string                `gorm:"size:42" json:"user_name"`                // 用户名
	NickName  string                `gorm:"size:42" json:"nick_name"`                // 昵称
	IP        string                `gorm:"size:20" json:"ip"`                       // ip
	Address   string                `gorm:"size:64" json:"address"`                  // 地址
	Device    string                `gorm:"size:36" json:"device"`                   // 登录设备
	Level     string                `gorm:"size:36" json:"level"`                    // 日志水平
	Content   string                `gorm:"size:64" json:"content"`                  // 日志内容
	LoginType model_type.SignStatus `gorm:"size:type=smallint(6)" json:"login_type"` // 登录方式
}

//	Level ->
//	Info ：一般信息（login、updatePwd、updateRole）
//	Debug：调试信息
//	Warn ：警告信息（用户名或密码错误、用户不存在）
//	Error：错误信息

//	Critical：严重错误信息（数据丢失）

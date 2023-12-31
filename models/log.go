package models

// Log 系统日志
type Log struct {
	MODEL
	UserName string `gorm:"size:42" json:"user_name"`  // 用户名
	NickName string `gorm:"size:42" json:"nick_name"`  // 昵称
	Email    string `gorm:"size:200" json:"email"`     // 邮箱
	IP       string `gorm:"size:20" json:"ip"`         // ip
	Address  string `gorm:"size:64" json:"address"`    // 地址
	Device   string `gorm:"size:36" json:"device"`     // 登录设备
	Level    string `gorm:"size:36" json:"level"`      // 日志水平
	Content  string `gorm:"mediumtext" json:"content"` // 日志内容
}

//	Level ->
//	Info ：一般信息（login、register）
//	Email ：邮箱信息（绑定邮箱、密码更新成功提醒、重置密码）
//	Warn ：警告信息（用户名或密码错误、用户不存在）

//	Error：错误信息（暂无）

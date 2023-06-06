package model_type

import "encoding/json"

type LoginType int

// 登录方式
const (
	LoginQQ       LoginType = 1 // QQ登录
	LoginUsername LoginType = 2 // 用户名登录
	LoginEmail    LoginType = 3 // 用户名登录
)

func (s LoginType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s LoginType) String() string {
	var str string
	switch s {
	case LoginQQ:
		str = "QQ"
	case LoginUsername:
		str = "用户名"
	case LoginEmail:
		str = "邮箱"
	default:
		str = "其他"
	}
	return str
}

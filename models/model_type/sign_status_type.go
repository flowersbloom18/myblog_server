package model_type

import "encoding/json"

type SignStatus int

const (
	SignQQ SignStatus = 1 // QQ注册
	Sign   SignStatus = 2 // 用户名注册
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case Sign:
		str = "用户名"
	default:
		str = "其他"
	}
	return str
}

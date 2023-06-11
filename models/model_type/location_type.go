package model_type

import "encoding/json"

// LocationType 存储位置类型
type LocationType int

const (
	Local LocationType = 1 // 本地
	QiNiu LocationType = 2 // 七牛云
)

func (s LocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s LocationType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	default:
		str = "其他"
	}
	return str
}

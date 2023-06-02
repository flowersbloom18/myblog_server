package model_type

import "encoding/json"

type Role int

const (
	PermissionAdmin   Role = 1 // 管理员
	PermissionUser    Role = 2 // 用户
	PermissionVisitor Role = 3 // 游客（具有管理员查看的权限,但有些数据依然无法查看）
)

// MarshalJSON 当转为json数据的时候自动转为对应的字体
func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	default:
		str = "其他"
	}
	return str
}

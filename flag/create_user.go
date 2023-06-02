package flag

import (
	"fmt"
	"myblog_server/global"
	model_type "myblog_server/models/model_type"
	"myblog_server/service/user_service"
)

func CreateUser(permissions string) {
	// 创建用户的逻辑
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码：")
	fmt.Scan(&rePassword)

	// 校验两次密码
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}
	role := model_type.PermissionUser
	if permissions == "admin" {
		role = model_type.PermissionAdmin
	}
	err := user_service.UserService{}.CreateUser(userName, nickName, password, role, email, "127.0.0.1", "未知设备")
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功!", userName)

}

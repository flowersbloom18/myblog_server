package user_service

import (
	"errors"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	ip2 "myblog_server/utils/ip"
	"myblog_server/utils/pwd"
)

const Avatar = "/uploads/avatar/favicon.png"

func (UserService) CreateUser(userName, nickName, password string, role model_type.Role, email string, ip string, device string) error {
	// 判断用户名是否存在
	var userModel models.User
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	// 对密码进行hash
	hashPwd := pwd.HashPwd(password)

	// 头像问题
	// 1. 默认头像
	// 2. 随机选择头像
	//Get
	address := ip2.GetAddressByIp(ip)
	// 入库
	err = global.DB.Create(&models.User{
		NickName:       nickName,
		UserName:       userName,
		Password:       hashPwd,
		Email:          email,
		Role:           role,
		Avatar:         Avatar,
		IP:             ip,
		Address:        address,
		RegisterOrigin: model_type.Sign,
		Device:         device,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

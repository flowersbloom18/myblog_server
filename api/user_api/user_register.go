package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/utils/device"
)

type UserRegister struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	NickName string `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	Password string `json:"password" binding:"required" msg:"请输入密码"`   // 密码
}

// UserRegisterView 创建用户
func (UserApi) UserRegisterView(c *gin.Context) {
	serviceApp := service.ServiceApp
	var cr UserRegister
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	device := device.GetLoginDevice(c)
	// 用户自己注册，权限为2固定
	err := serviceApp.UserService.CreateUser(cr.UserName, cr.NickName, cr.Password, 2, "", c.ClientIP(), device)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	global.DB.Create(&models.Log{
		UserName: cr.UserName,
		NickName: cr.NickName,
		IP:       c.ClientIP(),
		Device:   device,
		Level:    "Info",
		Content:  "注册成功",
	})

	response.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
	return
}

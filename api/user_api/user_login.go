package user_api

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/utils/device"
	ip2 "myblog_server/utils/ip"
	"myblog_server/utils/jwt"
	"myblog_server/utils/pwd"
	"strings"
)

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// UserLoginView 邮箱登录，返回token，用户信息需要从token中解码
func (UserApi) UserLoginView(c *gin.Context) {
	// 登录结果
	logContent := ""

	var cr LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 登录方式（qq、邮箱、用户名）
	var loginType model_type.LoginType
	str := strings.TrimSpace(cr.UserName)
	if strings.HasSuffix(str, ".com") {
		loginType = model_type.LoginEmail
	} else {
		loginType = model_type.LoginUsername
	}

	var userModel models.User
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		global.Log.Warn("用户名不存在")
		logContent := "登录中：用户名不存在！"
		global.DB.Create(&models.Log{
			Level:   "warn",
			Content: logContent,
		})
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")

		logContent := "用户名密码错误"
		global.DB.Create(&models.Log{
			UserName: userModel.UserName,
			NickName: userModel.NickName,
			IP:       userModel.IP,
			Address:  userModel.Address,
			Device:   userModel.Device,
			Level:    "warn",
			Content:  logContent,
		})
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwt.GenToken(jwt.PayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		Avatar:   userModel.Avatar,
	})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("token生成失败", c)
		return
	}

	// 获取ip和地址
	ip, addr := ip2.GetAddrByGin(c)
	global.Log.Info("\n 🥤userLogin63:ip= " + ip + "\taddr= " + addr)

	// 获取登录设备
	device := device.GetLoginDevice(c)

	// ⚠️登录之后需要修改用户的登录ip？addr？device？
	err = global.DB.Model(&userModel).Updates(map[string]interface{}{
		"ip":      ip,
		"address": addr,
		"device":  device,
	}).Error

	if err != nil {
		global.Log.Error(err)
		return
	}

	logContent = "通过（" + loginType.String() + "）登录成功"

	global.DB.Create(&models.Log{
		UserName: userModel.UserName,
		NickName: userModel.NickName,
		IP:       ip,
		Address:  addr,
		Device:   device,
		Level:    "info",
		Content:  logContent,
	})

	response.OkWithData(token, c)
}

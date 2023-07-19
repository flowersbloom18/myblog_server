package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/config"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
)

// SettingsInfoUpdateView 修改某一项的配置信息
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr) // 获取uri（唯一资源标识符号）

	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info) // 用于将HTTP请求的JSON数据绑定到结构体中
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Jwt = info
	case "upload":
		var info config.Upload
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Upload = info
	case "tianapi":
		var info config.TianApi
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.TianApi = info
	case "juhe":
		var info config.Juhe
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Juhe = info
	default:
		response.FailWithMessage("没有对应的配置信息", c)
		return
	}

	// 查看修改的用户信息并记录日志
	_claims, _ := c.Get("claims") // 当前登录用户解析后的信息
	claims := _claims.(*jwt.Claims)
	var user models.User
	err = global.DB.Take(&user, "id = ?", claims.UserID).Error
	if err != nil {
		global.Log.Warn("用户查询不到")
	}

	err = core.SetYaml()
	if err != nil {
		global.DB.Create(&models.Log{
			UserName: user.UserName,
			NickName: user.NickName,
			Email:    user.Email,
			IP:       user.IP,
			Address:  user.Address,
			Device:   user.Device,
			Level:    "Warn",
			Content:  "系统关键信息配置失败",
		})
		response.FailWithMessage("系统关键信息配置失败", c)
		return
	}

	global.DB.Create(&models.Log{
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		IP:       user.IP,
		Address:  user.Address,
		Device:   user.Device,
		Level:    "Info",
		Content:  "系统关键信息配置成功",
	})
	response.OkWithMessage("网站信息更新成功", c)
}

package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/config"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models/response"
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

	core.SetYaml()
	response.OkWith(c)
}

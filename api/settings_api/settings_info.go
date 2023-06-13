package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri

	// ShouldBindUri：可以将 URL 路径参数解析并绑定到 cr 结构体中的字段
	err := c.ShouldBindUri(&cr) //settingsInfo/:name
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site": // 站点配置
		response.OkWithData(global.Config.SiteInfo, c)
	case "email": // 邮箱配置
		response.OkWithData(global.Config.Email, c)
	case "qiniu": // 七牛云配置
		response.OkWithData(global.Config.QiNiu, c)
	case "jwt": // JWT配置
		response.OkWithData(global.Config.Jwt, c)
	case "upload": // 本地上传信息配置
		response.OkWithData(global.Config.Upload, c)
	case "tianapi": // tianapi配置
		response.OkWithData(global.Config.TianApi, c)
	case "juhe": // juhe配置
		response.OkWithData(global.Config.Juhe, c)
	default:
		response.FailWithMessage("没有对应的配置信息", c)
	}

}

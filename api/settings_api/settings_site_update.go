package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/config"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models/response"
)

// SettingsSiteUpdateView 编辑网站信息
func (SettingsApi) SettingsSiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = info
	core.SetYaml()
	response.OkWithMessage("网站信息更新成功", c)
}

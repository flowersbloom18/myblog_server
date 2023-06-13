package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
)

// SettingsSiteInfoView 显示网站信息
func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	response.OkWithData(global.Config.SiteInfo, c)
}

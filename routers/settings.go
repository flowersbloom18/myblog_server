package routers

import (
	"myblog_server/api"
)

func (router RouterGroup) Settings() {
	settingsApi := api.ApiGroupApp.SettingsApi
	// 前台获取站点信息
	router.GET("settings/site", settingsApi.SettingsSiteInfoView)
	// 修改前台站点信息
	router.PUT("settings/site", settingsApi.SettingsSiteUpdateView)
	// 后台获取站点详细信息
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	// 修改站点详细信息
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}

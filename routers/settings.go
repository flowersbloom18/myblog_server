package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Settings() {
	settingsApi := api.ApiGroupApp.SettingsApi

	// 前台访客
	// 前台获取站点信息
	router.GET("settings/site1", settingsApi.SettingsSiteInfoView)

	// 后台管理员
	// 后台获取站点详细信息
	router.GET("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoView)
	// 修改站点详细信息
	router.PUT("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoUpdateView)
}

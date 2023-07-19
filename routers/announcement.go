package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Announcement() {
	app := api.ApiGroupApp.AnnouncementApi
	// 获取
	router.GET("announcement", app.GetAnnouncementView)
	// 修改
	router.POST("announcement", middleware.JwtAdmin(), app.UpdateAnnouncementView)
}

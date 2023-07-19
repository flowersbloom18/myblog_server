package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) About() {
	app := api.ApiGroupApp.AboutAPI
	// 获取
	router.GET("about", app.GetAboutView)
	// 修改
	router.POST("about", middleware.JwtAdmin(), app.UpdateAboutView)
}

package routers

import (
	"myblog_server/api"
)

func (router RouterGroup) About() {
	app := api.ApiGroupApp.AboutAPI
	// 获取
	router.GET("about", app.GetAboutView)
	// 批量删除登录日志
	router.POST("about", app.UpdateAboutView)
}

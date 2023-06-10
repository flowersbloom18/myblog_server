package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Log() {
	app := api.ApiGroupApp.LoginApi
	// 获取系统日志
	router.GET("log", middleware.JwtAuth(), app.LogView)
	// 批量删除系统日志
	router.DELETE("log", middleware.JwtAdmin(), app.LogRemoveView)
}

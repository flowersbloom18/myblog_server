package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Log() {
	app := api.ApiGroupApp.LoginApi
	// 获取登录日志
	router.GET("log", middleware.JwtAuth(), app.LogView)
	// 批量删除登录日志
	router.DELETE("log", middleware.JwtAdmin(), app.LogRemoveView)
}

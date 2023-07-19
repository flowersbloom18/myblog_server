package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Music() {
	app := api.ApiGroupApp.MusicAPI
	// ⚠️权限给定！

	// 新增音乐
	router.POST("music", middleware.JwtAdmin(), app.MusicCreateView)
	// 查找-所有音乐
	router.GET("music", app.MusicListView)

	// 修改音乐
	router.PUT("music/:id", middleware.JwtAdmin(), app.MusicUpdateView)
	// 删除音乐
	router.DELETE("music", middleware.JwtAdmin(), app.MusicRemoveView)
}

package routers

import (
	"myblog_server/api"
)

func (router RouterGroup) Music() {
	app := api.ApiGroupApp.MusicAPI
	// ⚠️权限给定！

	// 新增音乐
	router.POST("music", app.MusicCreateView)
	// 查找-所有音乐
	router.GET("musics", app.MusicListView)

	// 修改音乐
	router.PUT("music/:id", app.MusicUpdateView)
	// 删除音乐
	router.DELETE("musics", app.MusicRemoveView)
}

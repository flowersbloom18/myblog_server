package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) FriendLink() {
	app := api.ApiGroupApp.FriendLinkAPI
	// ⚠️权限给定！

	// 新增友链
	router.POST("friendlink", middleware.JwtAdmin(), app.FriendLinkCreateView)
	// 查找-所有友链
	router.GET("friendlinks", app.FriendLinkListView)

	// 修改友链
	router.PUT("friendlink/:id", middleware.JwtAdmin(), app.FriendLinkUpdateView)
	// 删除友链
	router.DELETE("friendlinks", middleware.JwtAdmin(), app.FriendLinkRemoveView)
}

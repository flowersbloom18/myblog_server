package routers

import (
	"myblog_server/api"
)

func (router RouterGroup) FriendLink() {
	app := api.ApiGroupApp.FriendLinkAPI
	// ⚠️权限给定！

	// 新增友链
	router.POST("friendlink", app.FriendLinkCreateView)
	// 查找-所有友链
	router.GET("friendlinks", app.FriendLinkListView)

	// 修改友链
	router.PUT("friendlink/:id", app.FriendLinkUpdateView)
	// 删除友链
	router.DELETE("friendlinks", app.FriendLinkRemoveView)
}

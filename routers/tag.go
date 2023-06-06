package routers

import (
	"myblog_server/api"
)

func (router RouterGroup) Tag() {
	app := api.ApiGroupApp.TagApi
	// ⚠️权限给定！middleware.JwtAdmin(),

	// 新增标签
	router.POST("tag", app.TagCreateView)
	// 查找-所有标签
	router.GET("tags", app.TagListView)
	// 查找-某一个标签下所有博客
	router.GET("tag/:name", app.GetBlogsByTag)

	// 修改标签
	router.PUT("tag/:id", app.TagUpdateView)
	// 删除标签
	router.DELETE("tags", app.TagRemoveView)
}

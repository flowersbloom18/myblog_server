package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Blog() {
	app := api.ApiGroupApp.BlogApi
	// ⚠️权限给定！

	// 新增博客
	router.POST("blog", middleware.JwtAdmin(), app.BlogCreateView)
	// 查找-所有博客
	router.GET("blogs", app.BlogListView)
	// 查找-博客详情
	router.GET("blog/*link", app.BlogDetailView)

	// 修改博客
	router.PUT("blog/*link", app.BlogUpdateView)
	// 删除博客
	router.DELETE("blogs", app.BlogRemoveView)
}

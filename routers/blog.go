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

	// 编辑-博客内容-后台
	router.GET("blog/edit/:id", middleware.JwtAdmin(), app.BlogContentView)

	// 查找-博客详情-前台
	router.GET("blog/detail/*link", app.BlogDetailView)

	// 修改博客
	router.PUT("blog/:id", middleware.JwtAdmin(), app.BlogUpdateView)
	// 删除博客
	router.DELETE("blogs", middleware.JwtAdmin(), app.BlogRemoveView)

	// 新增点赞
	router.POST("blog/like/:id", app.BlogLikeView)

	// 获取博客总浏览量
	router.GET("blog/views", middleware.JwtAuth(), app.BlogAllPageView)
}

package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Category() {
	app := api.ApiGroupApp.CategoryApi
	// ⚠️权限给定！

	// 新增分类
	router.POST("category", middleware.JwtAdmin(), app.CategoryCreateView)
	// 查找-所有分类
	router.GET("categories", app.CategoryListView)
	// 查找-某一个分类下的所有博客
	router.GET("category/:name", app.GetBlogsByCategory)

	// 修改分类
	router.PUT("category/:id", middleware.JwtAdmin(), app.CategoryUpdateView)
	// 删除分类
	router.DELETE("categories", middleware.JwtAdmin(), app.CategoryRemoveView)
}

package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Collect() {
	api := api.ApiGroupApp.CollectApi
	// 点击一次收藏，再一次则取消
	router.POST("collect/:blogID", middleware.JwtAuth(), api.CollectCreateView)
	router.GET("collect", middleware.JwtAuth(), api.CollectListView)
	router.DELETE("collect", middleware.JwtAuth(), api.CollectRemoveView)
}

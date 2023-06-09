package routers

import (
	"myblog_server/api"
)

func (router RouterGroup) Info() {
	app := api.ApiGroupApp.InfoApi
	// 获取信息
	router.GET("info/:id", app.InfoListView)

	// 流程：
	// 请求数据->存入数据库->解析后返回。 定时更新信息
}

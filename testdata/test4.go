package main

import (
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models"
)

func main() {
	// 初始化yaml、连接mysql、初始化日志
	core.InitConf()
	global.DB = core.InitGorm()
	global.Log = core.InitLogger()
	db := global.DB

	var tag models.Tag
	if err := db.Preload("Blogs").First(&tag, "id=?", 7).Error; err != nil {
		// 处理获取博客记录失败的情况
		global.Log.Warn("err=", err)
	}

	global.Log.Info("⚠️tag = ", tag)
}

package main

import (
	"myblog_server/core"
	"myblog_server/global"
)

func main() {
	// 1、 读取配置文件
	core.InitConf()
	// 2、初始化日志
	global.Log = core.InitLogger()
	// 3、连接数据库
	global.DB = core.InitGorm()
}

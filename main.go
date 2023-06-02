package main

import (
	"myblog_server/core"
	"myblog_server/flag"
	"myblog_server/global"
	"myblog_server/routers"
	"myblog_server/utils/output"
)

func main() {
	// 1、读取Yaml配置文件
	core.InitConf()
	// 2、初始化日志
	global.Log = core.InitLogger()
	// 3、连接Mysql
	global.DB = core.InitGorm()
	// 4、连接Redis
	global.Redis = core.ConnectRedis()
	// 5、命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 6、初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()

	// 网站运行端口
	// 输出系统运行位置
	output.PrintSystem()

	// 路由运行端口
	err := router.Run(addr)
	if err != nil {
		global.Log.Warn(err)
	}
}

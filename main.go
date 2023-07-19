package main

import (
	"myblog_server/core"
	"myblog_server/flag"
	"myblog_server/global"
	"myblog_server/routers"
	"myblog_server/service"
	"myblog_server/utils/output"
	"time"
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

	// 7、创建定时任务调度器（定期更新信息数据）

	// 创建一个新的 Goroutine 来执行异步请求
	go func() {
		// 🥤系统执行前,执行一次操作,来更新一次信息数据
		service.ServiceApp.InfoService.UpdateInfoService()
	}()

	// 启动定时任务调度器 1次/30分钟更新热搜数据
	ticker := time.NewTicker(30 * time.Minute)
	go func() {
		for range ticker.C {
			// 执行定时任务的代码
			service.ServiceApp.InfoService.UpdateInfoService()
		}
	}()

	// 停止定时器
	defer ticker.Stop()

	// 8、网站运行端口,输出系统运行位置
	output.PrintSystem()

	// 路由运行端口
	err := router.Run(addr)
	if err != nil {
		global.Log.Warn(err)
	}
}

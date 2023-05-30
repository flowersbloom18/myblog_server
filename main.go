package main

import (
	"myblog_server/core"
	"myblog_server/flag"
	"myblog_server/global"
	"myblog_server/utils/output"
)

func main() {
	// 1、读取配置文件
	core.InitConf()
	// 2、初始化日志
	global.Log = core.InitLogger()
	// 3、连接Mysql
	global.DB = core.InitGorm()
	// 4、命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// 输出系统运行位置
	output.PrintSystem()

}

package flag

import (
	"myblog_server/global"
	"myblog_server/models"
)

func Makemigrations() {
	var err error
	// 生成表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.LogModel{},
			//go run main.go -db数据库迁移
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}

package flag

import (
	"myblog_server/global"
	"myblog_server/models"
)

func Makemigrations() {
	var err error

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.User{},       // 用户
			&models.Log{},        // 系统日志
			&models.Category{},   // 分类
			&models.Tag{},        // 标签
			&models.Blog{},       // 博客
			&models.Info{},       // 信息
			&models.About{},      // 关于
			&models.FriendLink{}, // 友链
			//go run main.go -db数据库迁移
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}

package flag

import (
	"myblog_server/global"
	"myblog_server/models"
)

func Makemigrations() {
	var err error

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.User{},         // 1、用户
			&models.Log{},          // 2、系统日志
			&models.Category{},     // 3、分类
			&models.Tag{},          // 4、标签（多对多一张表）
			&models.Blog{},         // 5、博客
			&models.Info{},         // 6、信息
			&models.About{},        // 7、关于
			&models.FriendLink{},   // 8、友链
			&models.Music{},        // 9、音乐
			&models.Collect{},      // 10、收藏（博客）
			&models.Comment{},      // 11、评论
			&models.CommentOpen{},  // 12、评论开关
			&models.Announcement{}, // 13、公告
			&models.Attachment{},   // 14、附件
			//go run main.go -db数据库迁移
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}

package music_service

import (
	"fmt"
	"gorm.io/gorm"
	"myblog_server/global"
	"myblog_server/models"
)

// name，author，url，cover，status，sort

func (MusicService) CreateMusic(name, author, url, cover string, status bool, sort int) error {
	db := global.DB

	// 检查音乐是否存在（若存在，则返回空。否则创建）
	var existingMusic models.Music
	err := db.Where("name = ?", name).First(&existingMusic).Error
	// 错误存在，且错误不为（找不到记录）才算做内部错误！
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("查找音乐失败: ", err.Error())
		return fmt.Errorf("查找音乐失败: %s", err.Error())
	}
	//global.Log.Info("existingMusic.ID=", existingMusic.ID)

	// 音乐存在判断，如果不存在则existingMusic.ID=0
	if existingMusic.ID != 0 {
		global.Log.Info("音乐已存在:", existingMusic.Name)
		return fmt.Errorf("'%s'音乐已存在", existingMusic.Name)
	}

	// 创建音乐
	Music := models.Music{
		Name:   name,
		Author: author,
		Url:    url,
		Cover:  cover,
		Status: status,
		Sort:   sort,
	}

	err = db.Create(&Music).Error
	if err != nil {
		global.Log.Error("创建音乐失败: ", err.Error())
		return fmt.Errorf("创建音乐失败: %s", err.Error())
	}

	global.Log.Info("音乐 '", Music.Name, " '创建成功")
	return nil
}

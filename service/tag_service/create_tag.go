package tag_service

import (
	"fmt"
	"gorm.io/gorm"
	"myblog_server/global"
	"myblog_server/models"
)

func (TagService) CreateTag(name, cover string) error {
	db := global.DB

	// 检查标签是否存在（若存在，则返回空。否则创建）
	var existingTag models.Tag
	// 如果存在，则err==nil
	err := db.Where("name = ?", name).First(&existingTag).Error
	// 错误存在，且错误不为（找不到记录）才算做内部错误！
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("查找标签失败，错误信息为：", err.Error())
		return fmt.Errorf("查找标签失败，错误信息为：%s", err.Error())
	}

	// 标签存在判断
	if existingTag.ID != 0 {
		global.Log.Info("标签已存在:", existingTag.Name)
		return fmt.Errorf("'%s'标签已存在", existingTag.Name)
	}

	// 创建标签
	tag := models.Tag{
		Name:  name,
		Cover: cover,
	}

	err = db.Create(&tag).Error
	if err != nil {
		global.Log.Error("创建标签失败: ", err.Error())
		return fmt.Errorf("创建标签失败: %s", err.Error())
	}

	global.Log.Info("标签 '", tag.Name, " '创建成功")
	return nil
}

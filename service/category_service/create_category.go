package category_service

import (
	"fmt"
	"gorm.io/gorm"
	"myblog_server/global"
	"myblog_server/models"
)

func (CategoryService) CreateCategory(name, cover string) error {
	db := global.DB

	// 检查分类是否存在（若存在，则返回空。否则创建）
	var existingCategory models.Category
	err := db.Where("name = ?", name).First(&existingCategory).Error
	// 错误存在，且错误不为（找不到记录）才算做内部错误！
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("查找分类失败: ", err.Error())
		return fmt.Errorf("查找分类失败: %s", err.Error())
	}
	global.Log.Info("existingCategory.ID=", existingCategory.ID)

	// 分类存在判断，如果不存在则existingCategory.ID=0
	if existingCategory.ID != 0 {
		global.Log.Info("分类已存在:", existingCategory.Name)
		return fmt.Errorf("'%s'分类已存在", existingCategory.Name)
	}

	// 创建分类
	category := models.Category{
		Name:  name,
		Cover: cover,
	}

	err = db.Create(&category).Error
	if err != nil {
		global.Log.Error("创建分类失败: ", err.Error())
		return fmt.Errorf("创建分类失败: %s", err.Error())
	}

	global.Log.Info("分类 '", category.Name, " '创建成功")
	return nil
}

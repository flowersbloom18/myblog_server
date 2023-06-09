package friendlink_service

import (
	"fmt"
	"gorm.io/gorm"
	"myblog_server/global"
	"myblog_server/models"
	"time"
)

// name，description，logo，url，isTop，topTime

func (FriendLinkService) CreateFriendLink(name, description, logo, url string, isTop bool) error {
	db := global.DB

	// 检查友链是否存在（若存在，则返回空。否则创建）
	var existingFriendLink models.FriendLink
	err := db.Where("name = ?", name).First(&existingFriendLink).Error
	// 错误存在，且错误不为（找不到记录）才算做内部错误！
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("查找友链失败: ", err.Error())
		return fmt.Errorf("查找友链失败: %s", err.Error())
	}
	//global.Log.Info("existingFriendLink.ID=", existingFriendLink.ID)

	// 友链存在判断，如果不存在则existingFriendLink.ID=0
	if existingFriendLink.ID != 0 {
		global.Log.Info("友链已存在:", existingFriendLink.Name)
		return fmt.Errorf("'%s'友链已存在", existingFriendLink.Name)
	}

	// 创建友链
	friendlink := models.FriendLink{
		Name:        name,
		Description: description,
		Logo:        logo,
		Url:         url,
		IsTop:       isTop,
		TopTime:     time.Now(),
	}

	err = db.Create(&friendlink).Error
	if err != nil {
		global.Log.Error("创建友链失败: ", err.Error())
		return fmt.Errorf("创建友链失败: %s", err.Error())
	}

	global.Log.Info("友链 '", friendlink.Name, " '创建成功")
	return nil
}

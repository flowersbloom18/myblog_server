package announcement_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

func (AnnouncementApi) GetAnnouncementView(c *gin.Context) {
	db := global.DB
	var announcement models.Announcement

	// 查询关于表的第一行数据
	result := db.First(&announcement)

	if result.Error != nil {
		// 处理查询错误
		global.Log.Info(fmt.Sprintf("查询公告信息失败:%s", result.Error))
	}

	// 判断查询结果是否为空
	if result.RowsAffected == 0 {
		// 创建公告表数据
		result := db.Create(&announcement)
		if result.Error != nil {
			// 处理创建错误
			fmt.Println("创建公告信息失败:", result.Error)
			return
		}
		announcement.Content = "目前没有公告。"
	}
	response.OkWithData(announcement, c)
}

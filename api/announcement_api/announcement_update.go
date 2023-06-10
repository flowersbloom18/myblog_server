package announcement_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

type AnnouncementRequest struct {
	Content string `json:"content"`
}

func (AnnouncementApi) UpdateAnnouncementView(c *gin.Context) {
	var cr AnnouncementRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c) // 错误，参数，上下文
		return
	}

	var announcement models.Announcement
	err = global.DB.First(&announcement).Error // 获取第一条数据
	if err != nil {
		global.Log.Info("更新失败")
		response.FailWithMessage("更新失败", c)
		return
	}

	err = global.DB.Model(announcement).Update("content", cr.Content).Error
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}

	response.OkWithMessage("公告信息更新成功", c)
	return
}

package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// CommentStatusView 获取评论开关信息
func (CommentApi) CommentStatusView(c *gin.Context) {
	// 1、首先判断开关数据是否存在， 如果不存在则创建数据
	count := global.DB.First(&models.CommentOpen{}).RowsAffected
	if count == 0 {
		err := global.DB.Create(&models.CommentOpen{
			IsOpen: true, // 默认建表,开启全局评论
		}).Error
		if err != nil {
			response.FailWithMessage("评论开启表创建失败", c)
			return
		}
	}

	// 创建之后返回数据
	var commentOpen models.CommentOpen
	err := global.DB.First(&commentOpen).Error
	if err != nil {
		response.FailWithMessage("数据获取错误", c)
	}

	response.OkWithData(commentOpen, c)
}

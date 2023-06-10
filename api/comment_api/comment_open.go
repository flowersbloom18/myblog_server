package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

type CommentOpenRequest struct {
	IsOpen bool `json:"is_open"`
}

// CommentOpenView 修改：一键【开启/关闭】评论
func (CommentApi) CommentOpenView(c *gin.Context) {
	var cr CommentOpenRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 2、修改数据
	var commentOpen models.CommentOpen
	err = global.DB.First(&commentOpen).Error
	if err != nil {
		response.FailWithMessage("数据获取错误", c)
		return
	}

	commentOpen.IsOpen = cr.IsOpen
	err = global.DB.Save(&commentOpen).Error
	if err != nil {
		response.FailWithMessage("数据更新错误", c)
		return
	}

	if cr.IsOpen {
		response.OkWithMessage("一键开启评论成功", c)
	} else {
		response.OkWithMessage("一键关闭评论成功", c)
	}
}
